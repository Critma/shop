package image_get_by_product

import (
	"context"
	"errors"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID: "get-image-by-product_v1",
		Method:      method,
		Path:        path,
		Summary:     "Get image by product",
		Description: "Get image for a product by product ID",
		Tags:        []string{"Images"},
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Image response",
				Content: map[string]*huma.MediaType{
					"application/octet-stream": {},
				},
			},
		},
	}, func(ctx context.Context, i *InputImageGetByProduct) (*OutputImageGetByProduct, error) {
		output, err := uc.GetImageByProduct(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("image not found", err)
			}
			return nil, render.ErrLeadInternalServer
		}

		output.ContentType = "application/octet-stream"
		return output, nil
	})
}
