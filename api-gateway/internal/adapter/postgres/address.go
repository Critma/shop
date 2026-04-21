package postgres

import (
	"context"
	"shopapi/internal/domain"
	"shopapi/pkg/zlog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) AddressCreate(ctx context.Context, tx pgx.Tx, address *domain.Address) (*domain.Address, error) {
	query := `INSERT INTO address (country, city, street) VALUES ($1, $2, $3) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	err := tx.QueryRow(ctx, query, address.Country, address.City, address.Street).Scan(&address.ID)
	if err != nil {
		zlog.PrintErr("pg.AddressCreate", address, err, "failed to create address")
		return nil, err
	}
	zlog.PrintInfo("pg.AddressCreate", address.ID, "address created")
	return address, err
}

func (pg *Postgres) AddressGetByID(ctx context.Context, addressID uuid.UUID) (*domain.Address, error) {
	query := `SELECT country, city, street FROM address WHERE id=$1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	adress := &domain.Address{
		ID: addressID,
	}
	err := pg.pool.QueryRow(ctx, query, addressID).Scan(&adress.Country, &adress.City, &adress.Street)
	if err != nil {
		zlog.PrintErr("pg.AddressGetByID", addressID, err, "failed to get address")
		return nil, err
	}
	zlog.PrintInfo("pg.AddressGetByID", addressID, "got address")
	return adress, nil
}

func (pg *Postgres) AddressGet(ctx context.Context, country, city, street string) (*domain.Address, error) {
	query := `SELECT id FROM address WHERE country ILIKE $1 and city ILIKE $2 and street ILIKE $3`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	adress := &domain.Address{
		Country: country,
		City:    city,
		Street:  street,
	}
	err := pg.pool.QueryRow(ctx, query, country, city, street).Scan(&adress.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		zlog.PrintErr("pg.AddressGet", adress, err, "failed to get address")
		return nil, err
	}
	zlog.PrintInfo("pg.AddressGet", adress, "got address")
	return adress, nil
}
