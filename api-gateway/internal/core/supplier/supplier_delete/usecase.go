package supplier_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	SupplierDelete(ctx context.Context, supplierID uuid.UUID) error
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

func (u *Usecase) SupplierDelete(ctx context.Context, input *InputSupplierDelete) error {
	err := u.store.SupplierDelete(ctx, input.SupplierID)
	if err != nil {
		return err
	}
	return nil
}
