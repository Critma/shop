package domain

import (
	"context"
)

type Storage interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, email string) (*User, error)
	UpdateUserPasswordHash(ctx context.Context, id, passwordHash string) error
}
