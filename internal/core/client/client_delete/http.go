package client_delete

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
		OperationID:   "delete-client_v1",
		Method:        method,
		Path:          path,
		Summary:       "Delete a client",
		Description:   "Delete a client by ID",
		Tags:          []string{"Clients"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputClientDelete) (*struct{}, error) {
		err := uc.ClientDelete(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, huma.Error404NotFound("client not found")
			}
			return nil, render.ErrLeadInternalServer
		}
		return nil, nil
	})
}
