package product_delete

import (
	"context"
	"net/http"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {

	huma.Register(api, huma.Operation{
		OperationID:   "delete-product_v1",
		Method:        method,
		Path:          path,
		Summary:       "Delete a product",
		Description:   "Delete a product by ID",
		Tags:          []string{"Products"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputProductDelete) (*struct{}, error) {
		err := uc.ProductDelete(ctx, i)
		if err != nil {
			if err == domain.ErrNotFound {
				return nil, huma.Error404NotFound("product not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		return nil, nil
	})
}
