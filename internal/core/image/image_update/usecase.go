package image_update

import (
	"context"
	"io"

	"github.com/google/uuid"
)

type Store interface {
	ImageUpdate(ctx context.Context, imageID uuid.UUID, imageData []byte) error
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) UpdateImage(ctx context.Context, input *InputImageUpdate) (*OutputImageUpdate, error) {
	formData := input.RawBody.Data()
	imageBytes, err := io.ReadAll(formData.Image)
	if err != nil {
		return nil, err
	}

	err = u.store.ImageUpdate(ctx, input.ImageID, imageBytes)
	if err != nil {
		return nil, err
	}

	output := &OutputImageUpdate{}
	output.Body.Success = true
	return output, nil
}
