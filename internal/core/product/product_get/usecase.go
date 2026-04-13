package product_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	ProductGet(ctx context.Context, productID uuid.UUID) (*domain.Product, error)
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetProduct(ctx context.Context, input *InputProductGet) (*OutputProductGet, error) {
	product, err := u.store.ProductGet(ctx, input.ProductID)
	if err != nil {
		return nil, err
	}
	output := &OutputProductGet{}
	output.Body.Product = product
	return output, nil
}
