package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	kafka_produce "github.com/critma/message-generator/internal/adapter/kafka_product"
	"github.com/critma/message-generator/internal/adapter/postgres"
)

type Config struct {
	// Host string `env:"SERVICE_HOST" envDefault:"localhost"`
	// Port string `env:"SERVICE_PORT" envDefault:"9091"`
	TickInterval time.Duration `env:"TICK_INTERVAL" envDefault:"10s"`

	Postgres postgres.Config `envPrefix:"POSTGRES_"`

	KafkaProduce kafka_produce.Config `envPrefix:"KAFKA_"`
}

func InitConfig() (config Config, err error) {
	err = env.Parse(&config)
	if err != nil {
		return
	}
	return
}
