package product_stock_decrease

import (
	"context"
	"shopapi/pkg/render"

	"github.com/google/uuid"
)

type Store interface {
	ProductDecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error
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

func (u *Usecase) ProductCreate(ctx context.Context, input *InputProductStockDecrease) (*OutputProductStockDecrease, error) {
	err := u.store.ProductDecreaseStock(ctx, input.Body.ProductID, input.Body.Quantity)
	if err != nil {
		return nil, render.ErrLeadInternalServer
	}
	output := &OutputProductStockDecrease{}
	return output, nil
}
