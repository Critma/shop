package grpc_controller

import (
	authv1 "github.com/critma/auth/gen/auth/v1"
	grpc_handlers "github.com/critma/auth/internal/core/handlers/grpc"
	"github.com/critma/auth/internal/core/usecase"
	"google.golang.org/grpc"
)

func RegisterAuthService(gRPCServer *grpc.Server, usecase *usecase.AuthUsecase) {
	authv1.RegisterAuthorizationServer(gRPCServer, grpc_handlers.NewAuthService(usecase))
}
