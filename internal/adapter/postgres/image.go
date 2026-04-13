package postgres

import (
	"context"
	"errors"
	"fmt"
	"shopapi/internal/domain"
	"shopapi/pkg/zlog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) ImageCreate(ctx context.Context, tx pgx.Tx, imageData []byte) (*domain.Image, error) {
	query := `INSERT INTO image(image) VALUES ($1) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	result := &domain.Image{Image: imageData}
	err := tx.QueryRow(ctx, query, imageData).Scan(&result.ID)
	if err != nil {
		zlog.PrintErr("pg.ImageCreate", imageData, err, "failed to create image")
		return nil, err
	}
	zlog.PrintInfo("pg.ImageCreate", result.ID, "image created")
	return result, nil
}

func (pg *Postgres) ImageUpdate(ctx context.Context, imageID uuid.UUID, imageData []byte) error {
	query := `UPDATE image SET image = $1 WHERE id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, imageData, imageID)
	if err != nil {
		zlog.PrintErr("pg.ImageUpdate", imageID, err, "failed to update image")
		return err
	}
	if tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ImageUpdate", imageID, err, "image not found")
		return domain.ErrNotFound
	}
	zlog.PrintInfo("pg.ImageUpdate", imageID, "image updated")
	return nil
}

func (pg *Postgres) ImageDelete(ctx context.Context, imageID uuid.UUID) error {
	query := `DELETE FROM image WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, imageID)
	if err != nil {
		zlog.PrintErr("pg.ImageDelete", imageID, err, "failed to delete image")
		return err
	}
	if tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ImageDelete", imageID, err, "image not found")
		return domain.ErrNotFound
	}
	zlog.PrintInfo("pg.ImageDelete", imageID, "image deleted")
	return nil
}

func (pg *Postgres) ImageGetByProduct(ctx context.Context, productID uuid.UUID) (*domain.Image, error) {
	query := `
	SELECT img.id, img.image
	 FROM image as img
	 JOIN product as p ON img.id = p.image_id
	 WHERE p.id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	image := &domain.Image{}
	err := pg.pool.QueryRow(ctx, query, productID).Scan(&image.ID, &image.Image)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		zlog.PrintErr("pg.ImageGetByProduct", productID, err, "failed to get image by product")
		return nil, err
	}
	zlog.PrintInfo("pg.ImageGetByProduct", image.ID, "image got by product")
	return image, nil
}

func (pg *Postgres) ImageGet(ctx context.Context, imageID uuid.UUID) (*domain.Image, error) {
	query := `SELECT image FROM image WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	image := &domain.Image{ID: imageID}
	err := pg.pool.QueryRow(ctx, query, imageID).Scan(&image.Image)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		zlog.PrintErr("pg.ImageGet", imageID, err, "failed to get image")
		return nil, err
	}
	zlog.PrintInfo("pg.ImageGet", imageID, "image got")
	return image, nil
}

func (pg *Postgres) ImageAssignToProduct(ctx context.Context, productID uuid.UUID, imageData domain.ImageData) (uuid.UUID, error) {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		zlog.PrintErr("pg.ImageAssignToProduct", fmt.Sprintf("productID: %v, imageData: %#v", productID, imageData), err, "failed to begin transaction")
		return uuid.Nil, err
	}
	image, err := pg.ImageCreate(ctx, tx, imageData)
	if err != nil {
		zlog.PrintErr("pg.ImageAssignToProduct", fmt.Sprintf("productID: %v, imageData: %#v", productID, imageData), err, "failed to create image")
		tx.Rollback(ctx)
		return uuid.Nil, err
	}
	if err := pg.ProductUpdateImage(ctx, tx, productID, image.ID); err != nil {
		zlog.PrintErr("pg.ImageAssignToProduct", fmt.Sprintf("productID: %v, imageID: %v", productID, image.ID), err, "failed to update product image")
		tx.Rollback(ctx)
		return uuid.Nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		zlog.PrintErr("pg.ImageAssignToProduct", fmt.Sprintf("productID: %v, imageID: %v", productID, image.ID), err, "failed to commit transaction")
		return uuid.Nil, err
	}
	return image.ID, nil
}
