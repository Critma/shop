package image_get_by_product

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	ImageGetByProduct(ctx context.Context, productID uuid.UUID) (*domain.Image, error)
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetImageByProduct(ctx context.Context, input *InputImageGetByProduct) (*OutputImageGetByProduct, error) {
	image, err := u.store.ImageGetByProduct(ctx, input.ProductID)
	if err != nil {
		return nil, err
	}

	output := &OutputImageGetByProduct{}
	output.Body = image.Image
	return output, nil
}
