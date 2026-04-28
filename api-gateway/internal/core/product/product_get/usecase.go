package product_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	ProductGet(ctx context.Context, productID uuid.UUID) (*domain.Product, error)
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

func (u *Usecase) GetProduct(ctx context.Context, input *InputProductGet) (*OutputProductGet, error) {
	product, err := u.store.ProductGet(ctx, input.ProductID)
	if err != nil {
		return nil, err
	}
	output := &OutputProductGet{}
	output.Body.Product = product
	return output, nil
}
