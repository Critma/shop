package auth_login

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, uc *useCase) {
	huma.Register(api, huma.Operation{
		OperationID:   "auth-login_v1",
		Method:        method,
		Path:          path,
		Summary:       "Login a user",
		Description:   "Login a user with email, password",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputLogin) (*OutputLogin, error) {
		output, err := uc.Login(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
