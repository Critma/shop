package supplier_update_address

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	SupplierChangeAddress(ctx context.Context, supplierID uuid.UUID, address *domain.Address) error
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

	return usecase
}

func (u *Usecase) UpdateSupplierAddress(ctx context.Context, input *InputUpdateSupplierAddress) (*OutputUpdateSupplierAddress, error) {
	err := u.store.SupplierChangeAddress(ctx, input.ID, toAddress(input))
	if err != nil {
		return nil, err
	}

	output := &OutputUpdateSupplierAddress{}
	output.Body.Success = true
	return output, nil
}
