package image_delete

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	ImageDelete(ctx context.Context, imageID uuid.UUID) error
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

func (u *Usecase) DeleteImage(ctx context.Context, input *InputImageDelete) error {
	err := u.store.ImageDelete(ctx, input.ImageID)
	if err != nil {
		return err
	}
	return nil
}
