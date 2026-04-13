package image_assign

import (
	"context"
	"io"
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type Store interface {
	ImageAssignToProduct(ctx context.Context, productID uuid.UUID, image domain.ImageData) (uuid.UUID, error)
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) ImageAssign(ctx context.Context, input *InputImageAssign) (*OutputImageAssign, error) {
	formData := input.RawBody.Data()
	data, err := io.ReadAll(formData.Image)
	if err != nil {
		return nil, err
	}
	imageUUID, err := u.store.ImageAssignToProduct(ctx, input.ProductID, data)
	if err != nil {
		return nil, err
	}
	output := &OutputImageAssign{}
	output.Body.ImageID = imageUUID
	return output, nil
}
