package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Config struct {
	User     string `env:"USER" envDefault:"admin"`
	Password string `env:"PASSWORD" envDefault:"admin"`
	DBName   string `env:"DBNAME" envDefault:"auth"`
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5440"`
}

type Postgres struct {
	pool *pgxpool.Pool
	log  *zap.Logger
}

var QueryTimeout = 5 * time.Second

func NewPostgres(ctx context.Context, cfg *Config, logger *zap.Logger) (*Postgres, error) {
	dbURL := "postgresql://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName
	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return &Postgres{pool: dbpool, log: logger}, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := p.pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}
	return nil
}

func (p *Postgres) Close() {
	p.pool.Close()
}
