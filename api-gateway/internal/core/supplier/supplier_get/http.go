package supplier_get

import (
	"context"
	"errors"
	"net/http"
	"shopapi/internal/domain"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string) {

	huma.Register(api, huma.Operation{
		OperationID:   "get-supplier_v1",
		Method:        method,
		Path:          path,
		Summary:       "Get a supplier",
		Description:   "Get a supplier by ID",
		Tags:          []string{"Suppliers"},
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputSupplierGet) (*OutputSupplierGet, error) {
		output, err := usecase.GetSupplier(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("supplier not found")
			}
			return nil, err
		}
		return output, nil
	})
}
