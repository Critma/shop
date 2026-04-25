package auth_login

import (
	"context"
	authv1 "shopapi/gen/auth/v1"
	"shopapi/internal/adapter/grpc_client"
	"shopapi/internal/domain"
)

type useCase struct {
	grpcClient *grpc_client.Client
}

func NewUsecase(grpcClient *grpc_client.Client) *useCase {

	return &useCase{
		grpcClient: grpcClient,
	}
}

func (u *useCase) Login(ctx context.Context, inp *InputLogin) (*OutputLogin, error) {
	if inp.Body.Email == "" || inp.Body.Password == "" {
		return nil, domain.ErrInvalidArgs
	}

	req := &authv1.LoginRequest{
		Email:    inp.Body.Email,
		Password: inp.Body.Password,
	}
	resp, err := u.grpcClient.AuthAPI.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}

	output := &OutputLogin{}
	output.Body.Token = resp.GetToken()
	return output, nil
}
