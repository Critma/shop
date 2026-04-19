package supplier_create

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {

	huma.Register(api, huma.Operation{
		OperationID:   "create-supplier_v1",
		Method:        method,
		Path:          path,
		Summary:       "Create a supplier",
		Description:   "Create a supplier with name, address and phone number",
		Tags:          []string{"Suppliers"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *InputSupplierCreate) (*OutputSupplierCreate, error) {
		output, err := uc.SupplierCreate(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}