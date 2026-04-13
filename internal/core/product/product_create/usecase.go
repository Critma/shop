package product_create

import (
	"context"
	"shopapi/internal/domain"
)

type Store interface {
	ProductCreate(ctx context.Context, product *domain.Product) (*domain.Product, error)
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) ProductCreate(ctx context.Context, input *InputProductCreate) (*OutputProductCreate, error) {
	createdProduct, err := u.store.ProductCreate(ctx, toProduct(input))
	if err != nil {
		return nil, err
	}
	output := &OutputProductCreate{}
	output.Body.ID = createdProduct.ID
	return output, nil
}
