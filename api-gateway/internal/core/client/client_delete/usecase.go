package client_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	ClientDelete(ctx context.Context, clientID uuid.UUID) error
}

type Usecase struct {
	store Store
}

var usecase *Usecase

func NewUsecase(store Store) *Usecase {
	uc := &Usecase{
		store: store,
	}

	usecase = uc

	return uc
}

func (u *Usecase) ClientDelete(ctx context.Context, input *InputClientDelete) error {
	err := u.store.ClientDelete(ctx, input.ClientID)
	if err != nil {
		return err
	}
	return nil
}
