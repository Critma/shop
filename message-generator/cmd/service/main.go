package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/critma/message-generator/config"
	kafka_produce "github.com/critma/message-generator/internal/adapter/kafka_product"
	"github.com/critma/message-generator/internal/adapter/postgres"
	"github.com/critma/message-generator/internal/controller/worker"
	"github.com/critma/message-generator/internal/usecase"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	config, err := config.InitConfig()
	if err != nil {
		log.Error("failed to init config", slog.Any("error", err))
		os.Exit(1)
	}
	log.Info("Config loaded", slog.Any("config", config))

	postgres, err := postgres.New(context.Background(), config.Postgres, log)
	if err != nil {
		log.Error("failed to connect to postgres", slog.Any("error", err))
		os.Exit(1)
	}
	defer postgres.Close()

	kafkaProduce := kafka_produce.NewProducer(config.KafkaProduce, log)
	defer kafkaProduce.Close()

	productUsecase := usecase.NewProductUsecase(postgres, log, kafkaProduce)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	workerServer := worker.NewWorkerServer(productUsecase, config.TickInterval, log)
	go func() {
		if err := workerServer.Run(context.Background()); err != nil {
			log.Error("failed to run worker", slog.Any("error", err))
			os.Exit(1)
		}
		os.Exit(0)
	}()
	<-sigChan
	workerServer.Stop()
	log.Info("WorkerServer stopped...")
}
