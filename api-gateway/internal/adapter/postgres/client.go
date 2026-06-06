package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shopapi/internal/domain"
	"shopapi/pkg/zlog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// without Address
func (pg *Postgres) ClientCreate(ctx context.Context, name, surname string, birthday time.Time, gender domain.Gender) (*domain.Client, error) {
	query := `INSERT INTO client(name, surname, birthday, gender) VALUES ($1, $2, $3, $4) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	client := &domain.Client{Name: name, Surname: surname, Birthday: birthday, Gender: gender}
	if err := pg.pool.QueryRow(ctx, query, name, surname, birthday, gender).Scan(&client.ID); err != nil {
		zlog.PrintErr("pg.ClientCreate", client, err, "failed to create client")
		return nil, err
	}
	zlog.PrintInfo("pg.ClientCreate", client.ID, "client created")
	return client, nil
}

func (pg *Postgres) ClientDelete(ctx context.Context, clientID uuid.UUID) error {
	query := `DELETE FROM client WHERE id=$1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, clientID)
	if err != nil {
		zlog.PrintErr("pg.ClientDelete", clientID, err, "failed to delete client")
		return err
	}
	if tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ClientDelete", clientID, err, "client not found")
		return errors.Join(err, domain.ErrNotFound)
	}
	zlog.PrintInfo("pg.ClientDelete", clientID, "client deleted")
	return nil
}

func (pg *Postgres) ClientsSearch(ctx context.Context, name, surname string) ([]*domain.Client, error) {
	query := `
		SELECT 
			client.id, 
			client.name, 
			client.surname, 
			client.birthday, 
			client.gender, 
			client.registration_date, 
			client.address_id,
			address.country, 
			address.city, 
			address.street
		FROM client
		LEFT JOIN address ON address.id = client.address_id
		WHERE 
			($1 = '' OR client.name ILIKE $1 || '%') AND
    		($2 = '' OR client.surname ILIKE $2 || '%')
		ORDER BY client.id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rows, err := pg.pool.Query(ctx, query, name, surname)
	if err != nil {
		zlog.PrintErr("pg.ClientsSearch", fmt.Sprintf("name=%s, surname=%s", name, surname), err, "failed to execute query")
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	clients := make([]*domain.Client, 0)
	for rows.Next() {
		client := &domain.Client{}
		var addressCountry, addressCity, addressStreet sql.NullString

		err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Surname,
			&client.Birthday,
			&client.Gender,
			&client.RegistrationDate,
			&client.AddressID,
			&addressCountry,
			&addressCity,
			&addressStreet,
		)
		if err != nil {
			zlog.PrintErr("pg.ClientsSearch", fmt.Sprintf("name=%s, surname=%s", name, surname), err, "failed to scan row")
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		if client.AddressID != nil {
			address := &domain.Address{ID: *client.AddressID}
			address.Country = addressCountry.String
			address.City = addressCity.String
			address.Street = addressStreet.String
			client.Address = address
		}
		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		zlog.PrintErr("pg.ClientsSearch", fmt.Sprintf("name=%s, surname=%s", name, surname), err, "row iteration error")
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	zlog.PrintInfo("pg.ClientsSearch", len(clients), "clients found")
	return clients, nil
}

// todo getclientByID and check existed address is already set
func (pg *Postgres) ClientChangeAddress(ctx context.Context, clientID uuid.UUID, address *domain.Address) error {
	existedAddress, err := pg.AddressGet(ctx, address.Country, address.City, address.Street)
	if err != nil {
		if !errors.Is(err, domain.ErrNotFound) {
			zlog.PrintErr("pg.ClientChangeAddress", fmt.Sprintf("clientID: %v, address: %#v", clientID, address), err, "failed to get address")
			return err
		}
	}

	tx, err := pg.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		zlog.PrintErr("pg.ClientChangeAddress", fmt.Sprintf("clientID: %v, address: %#v", clientID, address), err, "failed to begin transaction")
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

	query := `UPDATE client SET address_id=$1 WHERE id=$2`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()
	tag, err := tx.Exec(ctx, query, address.ID, clientID)
	if err != nil || tag.RowsAffected() != 1 {
		zlog.PrintErr("pg.ClientChangeAddress", fmt.Sprintf("clientID: %s, NewAddressID: %s", clientID, address.ID), err, "failed to change client address")
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	zlog.PrintInfo("pg.ClientChangeAddress", fmt.Sprintf("clientID: %s, NewAddressID: %s", clientID, address.ID), "client address updated")
	return nil
}

func (pg *Postgres) ClientGetList(ctx context.Context, params ListParams) ([]*domain.Client, error) {
	methodName := "pg.ClientGetList"
	query := `SELECT client.id, client.name, client.surname, birthday, gender, registration_date, address_id, address.country, address.city, address.street
	 FROM client
	 LEFT JOIN address on address.id = client.address_id
	 LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rows, err := pg.pool.Query(ctx, query, params.Limit, params.Offset)
	if err != nil {
		zlog.PrintErr(methodName, params, err, "failed to get client list")
		return nil, err
	}
	defer rows.Close()

	clients := make([]*domain.Client, 0)

	for rows.Next() {
		client := &domain.Client{}
		var addressCountry, addressCity, addressStreet sql.NullString
		err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Surname,
			&client.Birthday,
			&client.Gender,
			&client.RegistrationDate,
			&client.AddressID,
			&addressCountry,
			&addressCity,
			&addressStreet,
		)
		if err != nil {
			zlog.PrintErr(methodName, params, err, "failed to scan client")
			return nil, err
		}
		if client.AddressID != nil {
			address := &domain.Address{ID: *client.AddressID}
			address.Country = addressCountry.String
			address.City = addressCity.String
			address.Street = addressStreet.String
			client.Address = address
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		zlog.PrintErr(methodName, params, err, "failed in rows")
		return nil, err
	}

	zlog.PrintInfo(methodName, params, "client list got")
	return clients, nil
}
