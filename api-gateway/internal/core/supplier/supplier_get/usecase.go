package supplier_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	SupplierGet(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error)
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

func (u *Usecase) GetSupplier(ctx context.Context, input *InputSupplierGet) (*OutputSupplierGet, error) {
	supplier, err := u.store.SupplierGet(ctx, input.SupplierID)
	if err != nil {
		return nil, err
	}

	output := &OutputSupplierGet{}
	output.Body.Supplier = supplier
	return output, nil
}
