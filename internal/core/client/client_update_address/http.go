package client_update_address

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, store Store) {
	uc := newUseCase(store)

	huma.Register(api, huma.Operation{
		OperationID:   "update-address-client_v1",
		Method:        method,
		Path:          path,
		Summary:       "Update client address",
		Description:   "Update the address of a client by ID",
		Tags:          []string{"Clients"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputUpdateClientAddress) (*OutputUpdateClientAddress, error) {
		output, err := uc.UpdateClientAddress(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
