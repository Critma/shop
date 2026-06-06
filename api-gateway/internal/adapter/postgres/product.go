package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shopapi/internal/domain"
	"shopapi/pkg/zlog"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// product_id required
func (pg *Postgres) ProductCreate(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	query := `INSERT INTO product(name, category, price, available_stock, supplier_id, image_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	err := pg.pool.QueryRow(ctx, query, product.Name, product.Category, product.Price, product.AvailableStock, product.SupplierID, product.ImageID).Scan(&product.ID)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			zlog.PrintErr("pg.ProductCreate", product, err, "supplier or image not found")
			return nil, domain.ErrNotFound
		}
		zlog.PrintErr("pg.ProductCreate", product, err, "failed to create product")
		return nil, err
	}
	zlog.PrintInfo("pg.ProductCreate", product.ID, "product created")
	return product, nil
}

func (pg *Postgres) ProductDecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error {
	query := `UPDATE product SET available_stock = available_stock - $1 WHERE id = $2 AND available_stock >= $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, amount, productID)
	if err != nil || tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ProductDecreaseStock", fmt.Sprintf("productID: %s, amount: %d", productID, amount), err, "failed to decrease product stock")
		return err
	}
	zlog.PrintInfo("pg.ProductDecreaseStock", fmt.Sprintf("productID: %s, amount: %d", productID, amount), "product stock decreased")
	return nil
}

func (pg *Postgres) ProductGet(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	query := `
	SELECT product.name, product.category, product.price, product.available_stock, product.last_update_date, product.supplier_id, product.image_id,
	 supplier.id, supplier.name, supplier.address_id, supplier.phone_number,
	 address.country, address.city, address.street
	FROM product
	LEFT JOIN supplier ON supplier.id = product.supplier_id
	LEFT JOIN address ON address.id = supplier.address_id
	WHERE product.id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	product := &domain.Product{ID: productID}
	var supplierAddressCountry, supplierAddressCity, supplierAddressStreet sql.NullString
	err := pg.pool.QueryRow(ctx, query, productID).Scan(
		&product.Name, &product.Category, &product.Price, &product.AvailableStock, &product.LastUpdateDate, &product.SupplierID, &product.ImageID,
		&product.Supplier.ID, &product.Supplier.Name, &product.Supplier.AddressID, &product.Supplier.PhoneNumber,
		&supplierAddressCountry, &supplierAddressCity, &supplierAddressStreet,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			zlog.PrintErr("pg.ProductGet", productID, err, "product not found")
			return nil, domain.ErrNotFound
		}
		zlog.PrintErr("pg.ProductGet", productID, err, "failed to get product")
		return nil, err
	}
	if product.Supplier.AddressID != nil {
		product.Supplier.Address = &domain.Address{
			ID:      *product.Supplier.AddressID,
			Country: supplierAddressCountry.String,
			City:    supplierAddressCity.String,
			Street:  supplierAddressStreet.String,
		}
	}

	zlog.PrintInfo("pg.ProductGet", productID, "product got")
	return product, nil
}

// No inner info
func (pg *Postgres) ProductGetList(ctx context.Context, params ListParams) ([]*domain.Product, error) {
	methodName := "pg.ProductGetAll"
	query := `
	SELECT product.id, product.name, category, price, available_stock, last_update_date, supplier_id, image_id,
	   supplier.id, supplier.name, supplier.phone_number, supplier.address_id
	FROM product
	LEFT JOIN supplier ON supplier.id = product.supplier_id
	WHERE available_stock > 0
	LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rows, err := pg.pool.Query(ctx, query, params.Limit, params.Offset)
	if err != nil {
		zlog.PrintErr(methodName, nil, err, "failed to get product list")
		return nil, err
	}
	defer rows.Close()

	products := make([]*domain.Product, 0)

	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Price,
			&product.AvailableStock,
			&product.LastUpdateDate,
			&product.SupplierID,
			&product.ImageID,
			&product.Supplier.ID,
			&product.Supplier.Name,
			&product.Supplier.PhoneNumber,
			&product.Supplier.AddressID,
		)
		if err != nil {
			zlog.PrintErr(methodName, nil, err, "failed to scan product")
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		zlog.PrintErr(methodName, nil, err, "failed in rows")
		return nil, err
	}

	zlog.PrintInfo(methodName, nil, "product list got")
	return products, nil
}

func (pg *Postgres) ProductDelete(ctx context.Context, productID uuid.UUID) error {
	query := `DELETE FROM product WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, productID)
	if err != nil {
		zlog.PrintErr("pg.ProductDelete", productID, err, "failed to delete product")
		return err
	}
	if tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ProductDelete", productID, err, "product not found")
		return domain.ErrNotFound
	}
	zlog.PrintInfo("pg.ProductDelete", productID, "product deleted")
	return nil
}

func (pg *Postgres) ProductUpdateImage(ctx context.Context, tx pgx.Tx, productID, imageID uuid.UUID) error {
	query := `UPDATE product SET image_id=$1 WHERE id=$2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := tx.Exec(ctx, query, imageID, productID)
	if err != nil {
		zlog.PrintErr("pg.ProductUpdateImage", fmt.Sprintf("productId: %s, imageID: %s", productID, imageID), err, "failed to update product image")
		return err
	}

	if tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ProductUpdateImage", fmt.Sprintf("productId: %s, imageID: %s", productID, imageID), err, "product not found")
		return errors.Join(err, domain.ErrNotFound)
	}
	zlog.PrintInfo("pg.ProductUpdateImage", fmt.Sprintf("productId: %s, imageID: %s", productID, imageID), "product image updated")
	return nil
}
