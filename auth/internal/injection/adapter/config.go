package adapter

import (
	"context"

	"github.com/critma/auth/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideConfig(lc fx.Lifecycle, log *zap.Logger) (*config.Config, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ConfigToLog := *cfg
			if cfg.Token.Secret != "" {
				ConfigToLog.Token.Secret = "HIDDEN"
			}
			if cfg.Token.SecretGen {
				log.Warn("JWT-SECRET IS GENERATED, PLEASE USE YOUR OWN IN PROD")
			}
			log.Info("Config loaded", zap.Any("config", ConfigToLog))
			return nil
		},
	})
	return cfg, nil
}
