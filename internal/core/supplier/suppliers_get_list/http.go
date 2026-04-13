package suppliers_get_list

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, store Store) {
	uc := newUseCase(store)

	huma.Register(api, huma.Operation{
		OperationID:   "get-list-suppliers_v1",
		Method:        method,
		Path:          path,
		Summary:       "Get all suppliers",
		Description:   "Get all suppliers with optional pagination (limit and offset)",
		Tags:          []string{"Suppliers"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputGetAllSuppliers) (*OutputGetAllSuppliers, error) {
		output, err := uc.GetAllSuppliers(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}