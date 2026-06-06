package image_update

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Store interface {
	ImageUpdate(ctx context.Context, imageID uuid.UUID, imageData []byte) error
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

func (u *Usecase) UpdateImage(ctx context.Context, input *InputImageUpdate) (*OutputImageUpdate, error) {
	formData := input.RawBody.Data()
	imageBytes, err := io.ReadAll(formData.Image)
	if err != nil {
		return nil, err
	}

	err = u.store.ImageUpdate(ctx, input.ImageID, imageBytes)
	if err != nil {
		return nil, err
	}

	err = u.cache.DeleteImage(ctx, input.ImageID)
	if err != nil {
		log.Error().Err(err).Any("imageID", input.ImageID).Msg("failed to delete image from cache")
	}

	output := &OutputImageUpdate{}
	output.Body.Success = true
	return output, nil
}
