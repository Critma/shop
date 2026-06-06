package image_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Store interface {
	ImageGet(ctx context.Context, imageID uuid.UUID) (*domain.Image, error)
}

type Cache interface {
	GetImage(ctx context.Context, id uuid.UUID) (*domain.Image, error)
	SetImage(ctx context.Context, image *domain.Image) error
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

func (u *Usecase) GetImage(ctx context.Context, input *InputImageGet) (*OutputImageGet, error) {
	image, err := u.cache.GetImage(ctx, input.ImageID)
	if err != nil && err != domain.ErrNotFound {
		log.Error().Err(err).Msg("failed to get image from cache")
	} else if image != nil {
		log.Info().Str("imageID", input.ImageID.String()).Msg("image found in cache")
		output := &OutputImageGet{}
		output.Body = image.Image
		return output, nil
	}

	image, err = u.store.ImageGet(ctx, input.ImageID)
	if err != nil {
		return nil, err
	}

	// cached
	if err := u.cache.SetImage(ctx, image); err != nil {
		log.Error().Err(err).Any("imageID", image.ID).Msg("failed to set image to cache")
	}

	output := &OutputImageGet{}
	output.Body = image.Image
	return output, nil
}
