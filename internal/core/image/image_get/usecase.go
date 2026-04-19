package image_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	ImageGet(ctx context.Context, imageID uuid.UUID) (*domain.Image, error)
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetImage(ctx context.Context, input *InputImageGet) (*OutputImageGet, error) {
	image, err := u.store.ImageGet(ctx, input.ImageID)
	if err != nil {
		return nil, err
	}

	output := &OutputImageGet{}
	output.Body = image.Image
	return output, nil
}
