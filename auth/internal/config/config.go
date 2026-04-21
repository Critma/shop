package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/critma/auth/internal/adapter/postgres"
)

type JWTToken struct {
	TTL time.Duration `env:"TTL" envDefault:"1h"`
	// Secret string        `env:"SECRET"`
	Secret string `env:"SECRET" envDefault:"TOP_SECRET_TOP_SECRET_TOP_SECRET"`
}

type Config struct {
	ENV  string `env:"ENV" envDefault:"dev"`
	Host string `env:"AUTH_HOST" envDefault:"localhost"`
	Port int    `env:"AUTH_PORT" envDefault:"9090"`

	Postgres postgres.Config `envPrefix:"POSTGRES_"`

	Token JWTToken `envPrefix:"TOKEN_"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
