package supplier_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	SupplierGet(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error)
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetSupplier(ctx context.Context, input *InputSupplierGet) (*OutputSupplierGet, error) {
	supplier, err := u.store.SupplierGet(ctx, input.SupplierID)
	if err != nil {
		return nil, err
	}

	output := &OutputSupplierGet{}
	output.Body.Supplier = supplier
	return output, nil
}
