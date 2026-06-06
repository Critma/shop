package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const QUERY_TIMEOUT = 1 * time.Second

type Config struct {
	Addr     string        `env:"ADDR" envDefault:"localhost:6379"`
	Password string        `env:"PASSWORD" envDefault:""`
	CacheTTL time.Duration `env:"CACHE_TTL" envDefault:"10m"`
}

type Redis struct {
	cfg    Config
	client *redis.Client
}

func New(cfg Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error().Err(err).Msg("Unable to ping Redis")
		return nil, err
	}
	log.Info().Msgf("Ping to Redis success: %s", pong)
	return &Redis{
		cfg:    cfg,
		client: rdb,
	}, nil
}
