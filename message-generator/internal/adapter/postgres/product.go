package postgres

import (
	"context"
	"errors"
	"log/slog"
	"math/rand/v2"

	"github.com/critma/message-generator/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) ProductRandomStockPrice(ctx context.Context, productID uuid.UUID) error {
	query := `
	UPDATE product
	SET available_stock = $1, price = $2
	WHERE id = $3`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rndStock, rndPrice := rand.IntN(1000), rand.Float64()*1000
	tag, err := pg.pool.Exec(ctx, query, rndStock, rndPrice, productID)
	if err != nil || tag.RowsAffected() != 1 {
		pg.log.Error("pg.ProductDecreaseStock: failed to decrease product stock", slog.String("productID", productID.String()), slog.Int("stock", rndStock), slog.Float64("price", rndPrice), slog.Any("error", err))
		return err
	}
	return nil
}

func (pg *Postgres) ProductGetRandom(ctx context.Context) (*domain.Product, error) {
	id, err := pg.ProductGetRandomID(ctx)
	if err != nil {
		return nil, err
	}

	return pg.ProductGet(ctx, id)
}

func (pg *Postgres) ProductGet(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	query := `
	SELECT product.name, product.category, product.price, product.available_stock, product.last_update_date, product.supplier_id, product.image_id
	FROM product
	WHERE product.id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	product := &domain.Product{ID: productID}
	err := pg.pool.QueryRow(ctx, query, productID).Scan(
		&product.Name, &product.Category, &product.Price, &product.AvailableStock, &product.LastUpdateDate, &product.SupplierID, &product.ImageID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			pg.log.Error("pg.ProductGet : product not found", slog.String("productID", productID.String()), slog.Any("error", err))
			return nil, domain.ErrNotFound
		}
		pg.log.Error("pg.ProductGet : failed to get product", slog.String("productID", productID.String()), slog.Any("error", err))
		return nil, err
	}
	return product, nil
}

func (pg *Postgres) ProductGetRandomID(ctx context.Context) (uuid.UUID, error) {
	query := `SELECT id FROM product ORDER BY RANDOM() LIMIT 1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var productID uuid.UUID
	if err := pg.pool.QueryRow(ctx, query).Scan(&productID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			pg.log.Error("pg.ProductGetRandomID : no products found", slog.Any("error", err))
			return uuid.Nil, domain.ErrNotFound
		}
		pg.log.Error("pg.ProductGetRandomID : failed to get random product", slog.Any("error", err))
		return uuid.Nil, err
	}
	return productID, nil
}
