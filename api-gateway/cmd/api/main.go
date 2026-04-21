package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"shopapi/config"
	"shopapi/internal/adapter/postgres"
	"shopapi/internal/api"
	"time"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/rs/zerolog/log"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		return
	}

	// postgres, err := postgres.New(context.Background(), cfg.Postgres)
	postgres, err := postgres.New(context.Background(), postgres.NewConfig(cfg.PGUser, cfg.PGPassword, cfg.PGDBName, cfg.PGHost, cfg.PGPort))
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		os.Exit(1)
	}
	defer postgres.Close()

	app := config.App{
		Config: cfg,
		Store:  postgres,
	}

	// server
	cli := humacli.New(func(hooks humacli.Hooks, o *Options) {
		mux := api.GetRouter(app)
		srv := http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.ApiHost, o.Port),
			Handler: mux,
		}

		hooks.OnStart(func() {
			log.Info().Str("host", cfg.ApiHost).Int("port", o.Port).Msg("Server started")
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Error().Err(err).Msg("Failed to start server")
			}
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
		})
	})

	cli.Run()
}
