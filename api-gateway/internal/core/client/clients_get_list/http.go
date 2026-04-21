package clients_get_list

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID:   "get-list-clients_v1",
		Method:        method,
		Path:          path,
		Summary:       "Get all clients",
		Description:   "Get all clients with optional pagination (limit and offset)",
		Tags:          []string{"Clients"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputGetAllClients) (*OutputGetAllClients, error) {
		output, err := uc.GetAllClients(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
