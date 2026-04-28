package auth_header_middleware

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

func AuthMiddleware(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		authHeader := ctx.Header("Authorization")
		split := strings.Split(authHeader, " ")
		if authHeader == "" || len(split) != 2 || split[0] != "Bearer" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Authorization header must be in format 'Bearer <token>'")
			return
		}
		token := split[1]
		if token == "" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "token is empty")
			return
		}
		msg, err := usecase.TokenValidate(ctx.Context(), token)
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, msg)
			return
		}
		next(ctx)
	}
}
