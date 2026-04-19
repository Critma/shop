package supplier_create

import (
	"context"
	"shopapi/internal/domain"
)

type Store interface {
	SupplierCreate(ctx context.Context, supplier *domain.Supplier) (*domain.Supplier, error)
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) SupplierCreate(ctx context.Context, input *InputSupplierCreate) (*OutputSupplierCreate, error) {
	createdSupplier, err := u.store.SupplierCreate(ctx, toSupplier(input))
	if err != nil {
		return nil, err
	}

	output := &OutputSupplierCreate{}
	output.Body.ID = createdSupplier.ID
	return output, nil
}
