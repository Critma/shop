package injection

import (
	grpc_controller "github.com/critma/auth/internal/controller/grpc"
	"github.com/critma/auth/internal/core/usecase"
	"github.com/critma/auth/internal/domain"
	di_adapter "github.com/critma/auth/internal/injection/adapter"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func BuildGraph() *fx.App {
	ModuleUsecases := fx.Module("Usecases",
		fx.Provide(usecase.NewAuthUsecase),
	)

	ModulegRPCServices := fx.Module("gRPC services",
		fx.Invoke(grpc_controller.RegisterAuthService),
	)

	return fx.New(
		fx.Provide(
			zap.NewProduction,
			di_adapter.ProvideConfig,
			di_adapter.ProvidePostgres,
			di_adapter.ProvidegRPCServer,
		),
		ModuleUsecases,
		ModulegRPCServices,
		fx.Module("App",
			fx.Invoke(func(domain.Storage) {}),
			fx.Invoke(func(*grpc.Server) {}),
		),
	)
}
