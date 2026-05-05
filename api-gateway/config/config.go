package config

import (
	"shopapi/internal/adapter/kafka_consume"
	"shopapi/internal/adapter/postgres"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ApiHost string `env:"SERVICE_HOST" envDefault:"localhost"`

	Postgres postgres.Config `envPrefix:"POSTGRES_"`

	KafkaConsume kafka_consume.Config `envPrefix:"KAFKA_"`

	AuthHost string `env:"AUTH_HOST" envDefault:"localhost"`
	AuthPort string `env:"AUTH_PORT" envDefault:"9090"`
}

func InitConfig() (config Config, err error) {
	err = env.Parse(&config)
	if err != nil {
		log.Err(err).Msg("Failed to unmarshal config")
		return
	}
	log.Info().Any("config", config).Msg("Config loaded")
	return
}
