package config

import (
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	ApiHost string `mapstructure:"service_host"`

	// Postgres postgres.Config `mapstructure:"postgres"`
	PGUser     string `mapstructure:"POSTGRES_USER"    `
	PGPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PGDBName   string `mapstructure:"POSTGRES_DB" `
	PGHost     string `mapstructure:"POSTGRES_HOST"    `
	PGPort     string `mapstructure:"POSTGRES_PORT"    `
}

func InitConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../../config")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msg("Failed to read config")
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Err(err).Msg("Failed to unmarshal config")
		return
	}
	log.Info().Any("config", config).Msg("Config loaded")
	return
}
