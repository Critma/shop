package product_get

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
		OperationID:   "get-product_v1",
		Method:        method,
		Path:          path,
		Summary:       "Get a product",
		Description:   "Get a product by ID",
		Tags:          []string{"Products"},
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputProductGet) (*OutputProductGet, error) {
		output, err := usecase.GetProduct(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("product not found", err)
			}
			return nil, render.ErrLeadInternalServer
		}

		return output, nil
	})
}
