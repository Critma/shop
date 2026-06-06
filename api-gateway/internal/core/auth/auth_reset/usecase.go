package auth_reset

import (
	"context"
	authv1 "shopapi/gen/auth/v1"
	"shopapi/internal/adapter/grpc_client"
	"shopapi/internal/domain"
)

type Usecase struct {
	grpcClient *grpc_client.Client
}

var usecase *Usecase

func NewUsecase(grpcClient *grpc_client.Client) *Usecase {
	uc := &Usecase{
		grpcClient: grpcClient,
	}

	usecase = uc

	return uc
}

func (u *Usecase) Reset(ctx context.Context, input *InputReset) (*OutputReset, error) {
	if input.Body.Email == "" {
		return nil, domain.ErrInvalidArgs
	}

	req := &authv1.RecoverPasswordRequest{
		Email: input.Body.Email,
	}
	resp, err := u.grpcClient.AuthAPI.RecoverPassword(ctx, req)
	if err != nil {
		return nil, err
	}

	output := &OutputReset{}
	output.Body.Message = resp.GetMessage()
	return output, nil
}
