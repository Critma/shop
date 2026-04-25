package clients_search

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID:   "search-clients_v1",
		Method:        method,
		Path:          path,
		Summary:       "Search a clients",
		Description:   "Search a clients by name and surname",
		Tags:          []string{"Clients"},
		DefaultStatus: http.StatusOK,
		Security: []map[string][]string{
			{"bearer": {}},
		},
	}, func(ctx context.Context, i *InputClientsSearch) (*OutputClientsSearch, error) {
		output, err := uc.ClientsSearch(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
