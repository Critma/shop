package suppliers_get_list

import (
	"context"
	"shopapi/internal/adapter/postgres"
	"shopapi/internal/domain"
	"shopapi/pkg/render"
)

type Store interface {
	SupplierGetList(ctx context.Context, params postgres.ListParams) ([]*domain.Supplier, error)
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetAllSuppliers(ctx context.Context, input *InputGetAllSuppliers) (*OutputGetAllSuppliers, error) {
	var params postgres.ListParams
	// skip zero values
	if input.Limit != 0 {
		params.Limit = &input.Limit
	}
	if input.Offset != 0 {
		params.Offset = &input.Offset
	}

	suppliers, err := u.store.SupplierGetList(ctx, params)
	if err != nil {
		return nil, render.ErrLeadInternalServer
	}

	output := &OutputGetAllSuppliers{}
	output.Body.Suppliers = suppliers
	return output, nil
}
