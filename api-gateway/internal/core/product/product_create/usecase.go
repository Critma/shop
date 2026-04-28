package product_create

import (
	"context"
	"shopapi/internal/domain"
)

type Store interface {
	ProductCreate(ctx context.Context, product *domain.Product) (*domain.Product, error)
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

func (u *Usecase) ProductCreate(ctx context.Context, input *InputProductCreate) (*OutputProductCreate, error) {
	createdProduct, err := u.store.ProductCreate(ctx, toProduct(input))
	if err != nil {
		return nil, err
	}
	output := &OutputProductCreate{}
	output.Body.ID = createdProduct.ID
	return output, nil
}
