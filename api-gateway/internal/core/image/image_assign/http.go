package image_assign

import (
	"context"
	"errors"
	"net/http"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

const maxBody10MBytes = 1024 * 1024 * 10 // 10MB

func RegisterHTTPv1Handler(api huma.API, path, method string) {

	huma.Register(api, huma.Operation{
		OperationID:   "assign-image_v1",
		Method:        method,
		Path:          path,
		Summary:       "Assign image",
		Description:   "Assign image to a product",
		Tags:          []string{"Images"},
		DefaultStatus: http.StatusCreated,
		MaxBodyBytes:  maxBody10MBytes,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputImageAssign) (*OutputImageAssign, error) {
		output, err := usecase.ImageAssign(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("product or image not found", err)
			}
			return nil, render.ErrLeadInternalServer
		}
		return output, nil
	})
}
