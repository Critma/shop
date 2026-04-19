package supplier_update_address

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	SupplierChangeAddress(ctx context.Context, supplierID uuid.UUID, address *domain.Address) error
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) UpdateSupplierAddress(ctx context.Context, input *InputUpdateSupplierAddress) (*OutputUpdateSupplierAddress, error) {
	err := u.store.SupplierChangeAddress(ctx, input.ID, toAddress(input))
	if err != nil {
		return nil, err
	}

	output := &OutputUpdateSupplierAddress{}
	output.Body.Success = true
	return output, nil
}
