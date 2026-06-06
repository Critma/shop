package kafka_produce

import (
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	Addr  []string `env:"WRITER_ADDR" envDefault:"localhost:9092"`
	Topic string   `env:"WRITER_TOPIC" envDefault:"product-updated"`
}

type Producer struct {
	config Config
	writer *kafka.Writer
	log    *slog.Logger
}

func NewProducer(config Config, log *slog.Logger) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      config.Addr,
		Topic:        config.Topic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		MaxAttempts:  3,
	})
	writer.Compression = kafka.Snappy
	writer.AllowAutoTopicCreation = true

	return &Producer{
		config: config,
		writer: writer,
		log:    log,
	}
}

func (p *Producer) Close() {
	p.writer.Close()
}
