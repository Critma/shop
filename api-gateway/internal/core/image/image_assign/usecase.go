package image_assign

import (
	"context"
	"io"
	"shopapi/internal/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Store interface {
	ImageAssignToProduct(ctx context.Context, productID uuid.UUID, image domain.ImageData) (uuid.UUID, error)
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

func (u *Usecase) ImageAssign(ctx context.Context, input *InputImageAssign) (*OutputImageAssign, error) {
	formData := input.RawBody.Data()
	data, err := io.ReadAll(formData.Image)
	if err != nil {
		return nil, err
	}

	imageUUID, err := u.store.ImageAssignToProduct(ctx, input.ProductID, data)
	if err != nil {
		return nil, err
	}

	if err := u.cache.DeleteImage(ctx, imageUUID); err != nil {
		log.Error().Err(err).Msg("failed to delete image from cache")
	}

	output := &OutputImageAssign{}
	output.Body.ImageID = imageUUID
	return output, nil
}
