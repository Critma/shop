package config

import (
	"errors"
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

	AuthHost string `mapstructure:"AUTH_HOST"`
	AuthPort string `mapstructure:"AUTH_PORT"`
}

func InitConfig() (config Config, err error) {
	setDefaults()

	viper.AddConfigPath("./config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../../config")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Info().Msg("Config file not found, using defaults")
		} else {
			log.Error().Err(err).Msg("Failed to read config")
			return
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Err(err).Msg("Failed to unmarshal config")
		return
	}
	log.Info().Any("config", config).Msg("Config loaded")
	return
}

func setDefaults() {
	viper.SetDefault("POSTGRES_USER", "admin")
	viper.SetDefault("POSTGRES_PASSWORD", "admin")
	viper.SetDefault("POSTGRES_DB", "shopapi")
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_PORT", "5432")

	viper.SetDefault("PGADMIN_DEFAULT_EMAIL", "user@example.com")
	viper.SetDefault("PGADMIN_DEFAULT_PASSWORD", "securepassword")

	viper.SetDefault("SERVICE_HOST", "localhost")

	viper.SetDefault("AUTH_HOST", "host.docker.internal")
	viper.SetDefault("AUTH_PORT", "9090")
}
