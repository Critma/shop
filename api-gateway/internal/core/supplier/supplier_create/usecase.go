package supplier_create

import (
	"context"
	"shopapi/internal/domain"
)

type Store interface {
	SupplierCreate(ctx context.Context, supplier *domain.Supplier) (*domain.Supplier, error)
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

func (u *Usecase) SupplierCreate(ctx context.Context, input *InputSupplierCreate) (*OutputSupplierCreate, error) {
	createdSupplier, err := u.store.SupplierCreate(ctx, toSupplier(input))
	if err != nil {
		return nil, err
	}

	output := &OutputSupplierCreate{}
	output.Body.ID = createdSupplier.ID
	return output, nil
}
