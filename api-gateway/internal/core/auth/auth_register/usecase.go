package auth_register

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

func (u *useCase) Register(ctx context.Context, input *InputRegister) (*OutputRegister, error) {
	if input.Body.Email == "" || input.Body.Password == "" {
		return nil, domain.ErrInvalidArgs
	}

	req := &authv1.UserRequest{
		Email:    input.Body.Email,
		Name:     input.Body.Name,
		Surname:  input.Body.Surname,
		Phone:    input.Body.Phone,
		Password: input.Body.Password,
	}
	resp, err := u.grpcClient.AuthAPI.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	output := &OutputRegister{}
	output.Body.Token = resp.GetToken()
	return output, nil
}
