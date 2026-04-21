package products_get_list

import (
	"context"
	"shopapi/internal/adapter/postgres"
	"shopapi/internal/domain"
)

type Store interface {
	ProductGetList(ctx context.Context, params postgres.ListParams) ([]*domain.Product, error)
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetAllProducts(ctx context.Context, input *InputGetAllProducts) (*OutputGetAllProducts, error) {
	var params postgres.ListParams
	// skip zero values
	if input.Limit != 0 {
		params.Limit = &input.Limit
	}
	if input.Offset != 0 {
		params.Offset = &input.Offset
	}

	products, err := u.store.ProductGetList(ctx, params)
	if err != nil {
		return nil, err
	}

	output := &OutputGetAllProducts{}
	output.Body.Products = products
	return output, nil
}
