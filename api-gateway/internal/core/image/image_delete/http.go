package image_delete

import (
	"context"
	"errors"
	"net/http"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID:   "delete-image_v1",
		Method:        method,
		Path:          path,
		Summary:       "Delete an image",
		Description:   "Delete an image by ID",
		Tags:          []string{"Images"},
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputImageDelete) (*struct{}, error) {
		err := uc.DeleteImage(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("image not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		return nil, nil
	})
}
