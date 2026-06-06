package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shopapi/internal/domain"
	"shopapi/pkg/zlog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) SupplierCreate(ctx context.Context, supplier *domain.Supplier) (*domain.Supplier, error) {
	query := `INSERT INTO supplier(name, address_id, phone_number) VALUES ($1, $2, $3) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	err := pg.pool.QueryRow(ctx, query, supplier.Name, supplier.AddressID, supplier.PhoneNumber).Scan(&supplier.ID)
	if err != nil {
		zlog.PrintErr("pg.SupplierCreate", supplier, err, "failed to create supplier")
		return nil, err
	}
	zlog.PrintInfo("pg.SupplierCreate", supplier.ID, "supplier created")
	return supplier, nil
}

func (pg *Postgres) SupplierChangeAddress(ctx context.Context, supplierID uuid.UUID, address *domain.Address) error {
	existedAddress, err := pg.AddressGet(ctx, address.Country, address.City, address.Street)
	if err != nil {
		if !errors.Is(err, domain.ErrNotFound) {
			zlog.PrintErr("pg.SupplierChangeAddress", fmt.Sprintf("supplierID: %v, address: %#v", supplierID, address), err, "failed to get address")
			return err
		}
	}

	tx, err := pg.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		zlog.PrintErr("pg.SupplierChangeAddress", fmt.Sprintf("supplierID: %v, address: %#v", supplierID, address), err, "failed to begin transaction")
		return err
	}

	if existedAddress == nil {
		address, err = pg.AddressCreate(ctx, tx, address)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	} else {
		address = existedAddress
	}

	query := `UPDATE supplier SET address_id = $1 WHERE id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := tx.Exec(ctx, query, address.ID, supplierID)
	if err != nil || tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.SupplierChangeAddress", fmt.Sprintf("supplierID: %v, address: %#v", supplierID, address), err, "failed to change supplier address")
		tx.Rollback(ctx)
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	zlog.PrintInfo("pg.SupplierChangeAddress", fmt.Sprintf("supplierID: %v, address: %#v", supplierID, address), "supplier address updated")
	return nil
}

func (pg *Postgres) SupplierDelete(ctx context.Context, supplierID uuid.UUID) error {
	query := `DELETE FROM supplier WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, supplierID)
	if err != nil {
		zlog.PrintErr("pg.SupplierDelete", supplierID, err, "failed to delete supplier")
		return err
	}
	if tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.SupplierDelete", supplierID, err, "supplier not found")
		return domain.ErrNotFound
	}
	zlog.PrintInfo("pg.SupplierDelete", supplierID, "supplier deleted")
	return nil
}

func (pg *Postgres) SupplierGetList(ctx context.Context, params ListParams) ([]*domain.Supplier, error) {
	methodName := "pg.SupplierGetAll"
	query := `
	SELECT supplier.id, name, address_id, phone_number, address.country, address.city, address.street
	 FROM supplier
	 LEFT Join address on address.id = supplier.address_id
	 LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rows, err := pg.pool.Query(ctx, query, params.Limit, params.Offset)
	if err != nil {
		zlog.PrintErr(methodName, nil, err, "failed to get supplier list")
		return nil, err
	}
	defer rows.Close()

	suppliers := make([]*domain.Supplier, 0)

	for rows.Next() {
		supplier := &domain.Supplier{}
		var addressCountry, addressCity, addressStreet sql.NullString
		err := rows.Scan(
			&supplier.ID,
			&supplier.Name,
			&supplier.AddressID,
			&supplier.PhoneNumber,
			&addressCountry,
			&addressCity,
			&addressStreet,
		)
		if err != nil {
			zlog.PrintErr(methodName, nil, err, "failed to scan supplier")
			return nil, err
		}
		if supplier.AddressID != nil {
			address := &domain.Address{ID: *supplier.AddressID}
			address.Country = addressCountry.String
			address.City = addressCity.String
			address.Street = addressStreet.String
			supplier.Address = address
		}
		suppliers = append(suppliers, supplier)
	}

	if err := rows.Err(); err != nil {
		zlog.PrintErr(methodName, nil, err, "failed in rows")
		return nil, err
	}

	zlog.PrintInfo(methodName, nil, "supplier list got")
	return suppliers, nil
}

func (pg *Postgres) SupplierGet(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error) {
	query := `
	SELECT name, address_id, phone_number, address.country, address.city, address.street
	 FROM supplier
	 LEFT Join address on address.id = supplier.address_id
	 WHERE supplier.id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	supplier := &domain.Supplier{ID: supplierID}
	var addressCountry, addressCity, addressStreet sql.NullString
	err := pg.pool.QueryRow(ctx, query, supplierID).Scan(
		&supplier.Name, &supplier.AddressID, &supplier.PhoneNumber,
		&addressCountry, &addressCity, &addressStreet,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zlog.PrintInfo("pg.SupplierGet", supplierID, "supplier not found")
			return nil, domain.ErrNotFound
		}
		zlog.PrintErr("pg.SupplierGet", supplierID, err, "failed to get supplier")
		return nil, err
	}
	if supplier.AddressID != nil {
		address := &domain.Address{ID: *supplier.AddressID}
		address.Country = addressCountry.String
		address.City = addressCity.String
		address.Street = addressStreet.String
		supplier.Address = address
	}
	zlog.PrintInfo("pg.SupplierGet", supplierID, "supplier got")
	return supplier, nil
}
