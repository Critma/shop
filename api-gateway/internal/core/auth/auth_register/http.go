package auth_register

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shopapi/internal/domain"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID:   "auth-register_v1",
		Method:        method,
		Path:          path,
		Summary:       "Register a user",
		Description:   "Register a user with email, user_info, password",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *InputRegister) (*OutputRegister, error) {
		output, err := uc.Register(ctx, i)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidArgs) {
				return nil, huma.Error400BadRequest("invalid arguments")
			}
			return nil, fmt.Errorf("Internal error")
		}
		return output, nil
	})
}
