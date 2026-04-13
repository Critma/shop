package supplier_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	SupplierDelete(ctx context.Context, supplierID uuid.UUID) error
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) SupplierDelete(ctx context.Context, input *InputSupplierDelete) error {
	err := u.store.SupplierDelete(ctx, input.SupplierID)
	if err != nil {
		return err
	}
	return nil
}
