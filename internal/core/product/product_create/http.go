package product_create

import (
	"context"
	"net/http"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {

	huma.Register(api, huma.Operation{
		OperationID:   "create-product_v1",
		Method:        method,
		Path:          path,
		Summary:       "Create a product",
		Description:   "Create a product with optional imageID, required supplierID",
		Tags:          []string{"Products"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *InputProductCreate) (*OutputProductCreate, error) {
		output, err := uc.ProductCreate(ctx, i)
		if err != nil {
			if err == domain.ErrNotFound {
				return nil, huma.Error404NotFound("supplier or image not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		return output, nil
	})
}
