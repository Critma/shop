package image_update

import (
	"context"
	"errors"
	"net/http"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string) {
	huma.Register(api, huma.Operation{
		OperationID:   "update-image_v1",
		Method:        method,
		Path:          path,
		Summary:       "Update an image",
		Description:   "Update an image by ID with new image bytes data",
		Tags:          []string{"Images"},
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputImageUpdate) (*OutputImageUpdate, error) {
		output, err := usecase.UpdateImage(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("image not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		return output, nil
	})
}
