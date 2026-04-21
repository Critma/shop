package products_get_list

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {

	huma.Register(api, huma.Operation{
		OperationID:   "get-list-products_v1",
		Method:        method,
		Path:          path,
		Summary:       "Get all products",
		Description:   "Get all products with optional pagination (limit and offset)",
		Tags:          []string{"Products"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputGetAllProducts) (*OutputGetAllProducts, error) {
		output, err := uc.GetAllProducts(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}