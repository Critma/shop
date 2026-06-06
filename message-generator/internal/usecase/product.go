package usecase

import (
	"context"
	"log/slog"

	"github.com/critma/message-generator/internal/domain"
	"github.com/google/uuid"
)

type KafkaProduce interface {
	ProduceProductMessage(ctx context.Context, event domain.ProductChangeEvent) error
}

type Store interface {
	ProductGetRandom(ctx context.Context) (*domain.Product, error)
	ProductRandomStockPrice(ctx context.Context, productID uuid.UUID) error
	ProductGet(ctx context.Context, productID uuid.UUID) (*domain.Product, error)
}

type ProductUsecase struct {
	store        Store
	log          *slog.Logger
	kafkaProduce KafkaProduce
}

func NewProductUsecase(store Store, log *slog.Logger, kafkaProduce KafkaProduce) *ProductUsecase {
	return &ProductUsecase{
		store:        store,
		log:          log,
		kafkaProduce: kafkaProduce,
	}
}

func (u *ProductUsecase) ProductRandomStockPrice(ctx context.Context) error {
	product, err := u.store.ProductGetRandom(ctx)
	if err != nil {
		return err
	}

	if err := u.store.ProductRandomStockPrice(ctx, product.ID); err != nil {
		return err
	}

	updatedProduct, err := u.store.ProductGet(ctx, product.ID)
	if err != nil {
		return err
	}

	if err := u.kafkaProduce.ProduceProductMessage(ctx, domain.ProductChangeEvent{
		ProductID: product.ID.String(),
		OldPrice:  product.Price,
		NewPrice:  updatedProduct.Price,
		OldAmount: product.AvailableStock,
		NewAmount: updatedProduct.AvailableStock,
	}); err != nil {
		return err
	}

	return nil
}
