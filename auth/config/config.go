package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/critma/auth/internal/adapter/postgres"
	"github.com/critma/auth/pkg/jwt"
)

type JWTToken struct {
	TTL       time.Duration `env:"TTL" envDefault:"1h"`
	Secret    string        `json:"-"`
	SecretGen bool
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

	// get jwt secret from file
	cfg.Token.Secret, err = getSecret("jwt_secret")
	if err != nil {
		log.Printf("WARN: USED GENERATED JWT_SECRET, Please, write your own jwt secret!\n")
		jwt_gen_secret, err := jwt.GenerateJWTSecret(32)
		if err != nil {
			return nil, err
		}
		cfg.Token.Secret = jwt_gen_secret
		cfg.Token.SecretGen = true
	}

	return cfg, nil
}

func getSecret(name string) (string, error) {
	path := fmt.Sprintf("/run/secrets/%s", name)

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
