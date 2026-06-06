package image_get

import (
	"context"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string) {
	huma.Register(api, huma.Operation{
		OperationID: "get-image_v1",
		Method:      method,
		Path:        path,
		Summary:     "Get an image",
		Description: "Get an image by ID",
		Tags:        []string{"Images"},
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Image response",
				Content: map[string]*huma.MediaType{
					"application/octet-stream": {},
				},
			},
		},
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputImageGet) (*OutputImageGet, error) {
		output, err := usecase.GetImage(ctx, i)
		if err != nil {
			if err == domain.ErrNotFound {
				return nil, huma.Error404NotFound("image not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		output.ContentType = "application/octet-stream"
		return output, nil
	})
}
