package image_get

import (
	"context"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	ImageGet(ctx context.Context, imageID uuid.UUID) (*domain.Image, error)
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

func (u *Usecase) GetImage(ctx context.Context, input *InputImageGet) (*OutputImageGet, error) {
	image, err := u.store.ImageGet(ctx, input.ImageID)
	if err != nil {
		return nil, err
	}

	output := &OutputImageGet{}
	output.Body = image.Image
	return output, nil
}
