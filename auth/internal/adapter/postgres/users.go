package postgres

import (
	"context"
	"errors"

	"github.com/critma/auth/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

type LogUser struct {
	Email        string
	Name         string
	Surname      string
	Phone        string
	PasswordHash string
}

func newLogUser(user *domain.User) LogUser {
	return LogUser{
		Email:        user.Email,
		Name:         user.Name,
		Surname:      user.Surname,
		Phone:        user.Phone,
		PasswordHash: string(user.Password.Hash),
	}
}

// password_hash needed
func (pg *Postgres) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users(email, name, surname, phone, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if err := pg.pool.QueryRow(ctx, query, user.Email, user.Name, user.Surname, user.Phone, user.Password.Hash).Scan(&user.ID, &user.Created, &user.Updated); err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				return nil, domain.ErrAlreadyExists
			}
		}
		pg.log.Error("pg_CreateUser: failed to create user", zap.Any("input", newLogUser(user)), zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (pg *Postgres) GetUser(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, name, surname, phone, password_hash, created_at, updated_at FROM users WHERE email = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	user := &domain.User{}
	if err := pg.pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.Phone, &user.Password.Hash, &user.Created, &user.Updated); err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		pg.log.Error("pg_Login: failed to get user", zap.String("email", email), zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (pg *Postgres) UpdateUserPasswordHash(ctx context.Context, id, passwordHash string) error {
	query := `UPDATE users SET password_hash = $1 WHERE id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	tag, err := pg.pool.Exec(ctx, query, passwordHash, id)
	if err != nil {
		pg.log.Error("pg_UpdateUserPassword: failed to update user password", zap.String("id", id), zap.String("password_hash", passwordHash), zap.Error(err))
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}
