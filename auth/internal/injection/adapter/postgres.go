package adapter

import (
	"context"
	"fmt"

	"github.com/critma/auth/config"
	"github.com/critma/auth/internal/adapter/postgres"
	"github.com/critma/auth/internal/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvidePostgres(cfg *config.Config, lc fx.Lifecycle, log *zap.Logger) (domain.Storage, error) {
	pg, err := postgres.NewPostgres(context.Background(), &cfg.Postgres, log)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := pg.Ping(ctx); err != nil {
				return fmt.Errorf("%s:%s: %w", cfg.Postgres.Host, cfg.Postgres.Port, err)
			}
			log.Info("Ping to Postgres success")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Close postgres")
			pg.Close()
			return nil
		},
	})
	return pg, nil
}
