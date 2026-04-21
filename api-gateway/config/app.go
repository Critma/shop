package config

import "shopapi/internal/adapter/postgres"

type App struct {
	Config Config
	Store  *postgres.Postgres
}
