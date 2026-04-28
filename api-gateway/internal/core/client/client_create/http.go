package client_create

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-client_v1",
		Method:        method,
		Path:          path,
		Summary:       "Create a client",
		Description:   "Create a client with name, username, birthday and gender",
		Tags:          []string{"Clients"},
		DefaultStatus: http.StatusCreated,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputClientCreate) (*OutputClientCreate, error) {
		output, err := usecase.ClientCreate(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
