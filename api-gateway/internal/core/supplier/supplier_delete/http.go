package supplier_delete

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
		OperationID:   "delete-supplier_v1",
		Method:        method,
		Path:          path,
		Summary:       "Delete a supplier",
		Description:   "Delete a supplier by ID",
		Tags:          []string{"Suppliers"},
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputSupplierDelete) (*struct{}, error) {
		err := uc.SupplierDelete(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("supplier not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		return nil, nil
	})
}
