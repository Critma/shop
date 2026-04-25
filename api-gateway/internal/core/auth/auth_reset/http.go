package auth_reset

import (
	"context"
	"errors"
	"net/http"
	"shopapi/internal/domain"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID:   "auth-reset_v1",
		Method:        method,
		Path:          path,
		Summary:       "Reset a user password",
		Description:   "Reset a user password from email",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusAccepted,
	}, func(ctx context.Context, i *InputReset) (*OutputReset, error) {
		output, err := uc.Reset(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidArgs) {
				return nil, huma.Error400BadRequest("invalid arguments")
			}
			return nil, err
		}
		return output, nil
	})
}
