package auth_header_middleware

import (
	"context"
	authv1 "shopapi/gen/auth/v1"
	"shopapi/internal/adapter/grpc_client"
)

var usecase *Usecase

type Usecase struct {
	grpcClient *grpc_client.Client
}

func NewUsecase(grpcClient *grpc_client.Client) *Usecase {
	uc := &Usecase{
		grpcClient: grpcClient,
	}

	usecase = uc

	return uc
}

func (u *Usecase) TokenValidate(ctx context.Context, token string) (string, error) {
	resp, err := u.grpcClient.AuthAPI.ValidateToken(ctx, &authv1.TokenRequest{
		Token: token,
	})
	if err != nil {
		return "", err
	}

	return resp.GetMessage(), nil
}
