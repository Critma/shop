package adapter

import (
	"context"

	"github.com/critma/auth/config"
	grpc_controller "github.com/critma/auth/internal/controller/grpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ProvidegRPCServer(lc fx.Lifecycle, cfg *config.Config, log *zap.Logger) *grpc.Server {
	srv := grpc_controller.NewgRPCServer(log)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go grpc_controller.Start(cfg, srv, log)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpc_controller.Stop(srv, log)
			return nil
		},
	})
	return srv
}
