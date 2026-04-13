package image_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	ImageDelete(ctx context.Context, imageID uuid.UUID) error
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) DeleteImage(ctx context.Context, input *InputImageDelete) error {
	err := u.store.ImageDelete(ctx, input.ImageID)
	if err != nil {
		return err
	}
	return nil
}
