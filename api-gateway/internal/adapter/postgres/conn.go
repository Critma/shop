package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Config struct {
	User     string `mapstructure:"user"    `
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db" `
	Host     string `mapstructure:"host"    `
	Port     string `mapstructure:"port"    `
}

func NewConfig(user, password, dbname, host, port string) Config {
	return Config{
		User:     user,
		Password: password,
		DBName:   dbname,
		Host:     host,
		Port:     port,
	}
}

type Postgres struct {
	pool *pgxpool.Pool
}

var QueryTimeout = 5 * time.Second

func New(ctx context.Context, cfg Config) (*Postgres, error) {
	dbURL := "postgresql://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName
	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	log.Info().Msg("Connection to Postgres created")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}
	log.Info().Msg("Ping to Postgres success")
	return &Postgres{pool: dbpool}, nil
}

func (p *Postgres) Close() {
	p.pool.Close()
}
