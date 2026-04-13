package supplier_update_address

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, store Store) {
	uc := newUseCase(store)

	huma.Register(api, huma.Operation{
		OperationID:   "update-address-supplier_v1",
		Method:        method,
		Path:          path,
		Summary:       "Update supplier address",
		Description:   "Update the address of a supplier",
		Tags:          []string{"Suppliers"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputUpdateSupplierAddress) (*OutputUpdateSupplierAddress, error) {
		output, err := uc.UpdateSupplierAddress(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
