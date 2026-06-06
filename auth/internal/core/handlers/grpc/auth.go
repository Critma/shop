package grpc_handlers

import (
	"context"
	"errors"

	authv1 "github.com/critma/auth/gen/auth/v1"
	"github.com/critma/auth/internal/core/usecase"
	"github.com/critma/auth/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	authv1.UnimplementedAuthorizationServer
	usecase *usecase.AuthUsecase
}

func NewAuthService(usecase *usecase.AuthUsecase) *authService {
	return &authService{
		usecase: usecase,
	}
}

func (s *authService) CreateUser(ctx context.Context, req *authv1.UserRequest) (*authv1.TokenReply, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Surname == "" {
		return nil, status.Error(codes.InvalidArgument, "surname is required")
	}
	if req.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, "phone is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	user := domain.NewUser(req.Email, req.Name, req.Surname, req.Phone, req.Password)
	tokenString, err := s.usecase.UserRegister(ctx, user)
	if err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.InvalidArgument, "failed to register user")
	}

	return &authv1.TokenReply{Token: tokenString}, nil
}

func (s *authService) LoginUser(ctx context.Context, req *authv1.LoginRequest) (*authv1.TokenReply, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	tokenString, err := s.usecase.UserLogin(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return nil, status.Error(codes.PermissionDenied, "invalid credentials")
		}
		return nil, status.Error(codes.Internal, "failed to login user")
	}
	return &authv1.TokenReply{Token: tokenString}, nil
}

func (s *authService) ChangePassword(ctx context.Context, req *authv1.ChangePasswordRequest) (*authv1.MessageReply, error) {
	if req.LoginCredentials.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.LoginCredentials.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "old password is required")
	}
	if req.NewPassword == "" {
		return nil, status.Error(codes.InvalidArgument, "new password is required")
	}

	if err := s.usecase.PasswordChange(ctx, req.LoginCredentials.Email, req.LoginCredentials.Password, req.NewPassword); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return nil, status.Error(codes.PermissionDenied, "invalid credentials")
		}
		return nil, status.Error(codes.Internal, "failed to change password")
	}
	return &authv1.MessageReply{
		Message: "Password changed successfully",
	}, nil
}

func (s *authService) RecoverPassword(ctx context.Context, req *authv1.RecoverPasswordRequest) (*authv1.MessageReply, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if err := s.usecase.PasswordRecovery(ctx, req.Email); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to recover password")
	}
	return &authv1.MessageReply{Message: "Password recovery email sent successfully"}, nil
}

func (s *authService) ValidateToken(ctx context.Context, req *authv1.TokenRequest) (*authv1.MessageReply, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	ok, err := s.usecase.TokenValidate(req.Token)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "token validate fail: "+err.Error())
	}

	if ok {
		return &authv1.MessageReply{
			Message: "Token is valid",
		}, nil
	}
	return nil, status.Error(codes.PermissionDenied, "token is invalid")
}
