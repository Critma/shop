package kafka_produce

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/critma/message-generator/internal/domain"
	"github.com/segmentio/kafka-go"
)

func (p *Producer) ProduceProductMessage(ctx context.Context, event domain.ProductChangeEvent) error {
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.ProductID),
		Value: value,
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		p.log.Error("Error sending to kafka", slog.Any("error", err))
		return err
	}

	return nil
}
