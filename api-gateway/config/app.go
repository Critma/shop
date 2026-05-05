package config

import (
	"shopapi/internal/adapter/grpc_client"
	"shopapi/internal/adapter/kafka_consume"
	"shopapi/internal/adapter/postgres"
)

type App struct {
	Config        Config
	Store         *postgres.Postgres
	GrpcClient    *grpc_client.Client
	KafkaConsumer *kafka_consume.Consumer
}
