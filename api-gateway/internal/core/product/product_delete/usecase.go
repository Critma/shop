package product_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	ProductDelete(ctx context.Context, productID uuid.UUID) error
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

func (u *Usecase) ProductDelete(ctx context.Context, input *InputProductDelete) error {
	err := u.store.ProductDelete(ctx, input.ProductID)
	if err != nil {
		return err
	}
	return nil
}
