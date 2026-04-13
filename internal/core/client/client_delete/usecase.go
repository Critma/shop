package client_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	ClientDelete(ctx context.Context, clientID uuid.UUID) error
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) ClientDelete(ctx context.Context, input *InputClientDelete) error {
	err := u.store.ClientDelete(ctx, input.ClientID)
	if err != nil {
		return err
	}
	return nil
}
