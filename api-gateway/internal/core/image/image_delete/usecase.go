package image_delete

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Store interface {
	ImageDelete(ctx context.Context, imageID uuid.UUID) error
}

type Cache interface {
	DeleteImage(ctx context.Context, id uuid.UUID) error
}

type Usecase struct {
	store Store
	cache Cache
}

var usecase *Usecase

func NewUsecase(store Store, cache Cache) *Usecase {
	uc := &Usecase{
		store: store,
		cache: cache,
	}

	usecase = uc

	return uc
}

func (u *Usecase) DeleteImage(ctx context.Context, input *InputImageDelete) error {
	err := u.store.ImageDelete(ctx, input.ImageID)
	if err != nil {
		return err
	}

	if err := u.cache.DeleteImage(ctx, input.ImageID); err != nil {
		log.Error().Err(err).Msg("failed to delete image from redis")
	}

	return nil
}
