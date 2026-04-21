package product_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	ProductDelete(ctx context.Context, productID uuid.UUID) error
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) ProductDelete(ctx context.Context, input *InputProductDelete) error {
	err := u.store.ProductDelete(ctx, input.ProductID)
	if err != nil {
		return err
	}
	return nil
}
