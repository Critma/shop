package kafka_consume

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"shopapi/internal/adapter/websocket"
	"shopapi/internal/domain"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Config struct {
	Addr    []string `env:"READER_ADDR" envDefault:"localhost:9094"`
	Topic   string   `env:"READER_TOPIC" envDefault:"product-updated"`
	GroupID string   `env:"READER_GROUP_ID" envDefault:"api-gateway-local"`
}

type Consumer struct {
	reader *kafka.Reader
	kChan  chan kafka.Message
	sse    *sse
	ws     *websocket.Websocket
}

func NewConsumer(config Config, ws *websocket.Websocket) (*Consumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        config.Addr,
		Topic:          config.Topic,
		GroupID:        config.GroupID,
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
		MaxWait:        10 * time.Second,
		ReadBackoffMin: 100 * time.Millisecond,
		ReadBackoffMax: 1 * time.Second,
	})

	if !isKafkaAvailable(config.Addr, 10*time.Second) {
		log.Error().Msg("Kafka is not available")
		return nil, errors.New("kafka is not available")
	}

	return &Consumer{
		reader: reader,
		kChan:  make(chan kafka.Message),
		sse:    newSee(),
		ws:     ws,
	}, nil
}

func isKafkaAvailable(brokers []string, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	d := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	conn, err := d.DialContext(ctx, "tcp", brokers[0])
	if err != nil {
		return false
	}
	defer conn.Close()

	_, err = conn.ReadPartitions("any-topic-name")
	return err == nil
}

type sse struct {
	sseClients map[chan domain.ProductChangeEvent]bool
	mutex      sync.RWMutex
}

func newSee() *sse {
	return &sse{
		sseClients: make(map[chan domain.ProductChangeEvent]bool),
		mutex:      sync.RWMutex{},
	}
}

func (c *Consumer) StartPoiling(ctx context.Context) {
	const timeout = 120 * time.Second
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("kafka consumer context cancelled, stopping")
			return
		default:
			fetchCtx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
			msg, err := c.reader.FetchMessage(fetchCtx)
			cancel()

			if err != nil {
				if ctx.Err() != nil || errors.Is(err, io.EOF) {
					log.Info().Msg("context canceled or get EOF, stopping consumer")
					return
				}

				var deadlineExceeded interface{ Timeout() bool }
				if errors.As(err, &deadlineExceeded) && deadlineExceeded.Timeout() {
					log.Debug().Any("wait time, ms", timeout).Msg("kafka fetch timeout: no messages available")
				} else {
					log.Error().Err(err).Msg("failed to fetch message from kafka")
				}

				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Second):
					continue
				}
			}

			var event domain.ProductChangeEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Error().Any("error", err).Msg("failed to unmarshal kafka message")
				if err := c.reader.CommitMessages(ctx, msg); err != nil {
					log.Error().Any("error", err).Msg("failed to commit kafka message")
				}
				continue
			}
			event.ProductID = string(msg.Key)

			log.Info().Any("event", event).Msg("received event from kafka")

			if c.ws != nil {
				c.ws.Broadcast(map[string]any{
					"productID": event.ProductID,
					"oldPrice":  event.OldPrice,
					"newPrice":  event.NewPrice,
					"OldStock":  event.OldAmount,
					"NewStock":  event.NewAmount,
				})
			}

			if c.sse != nil {
				c.sse.broadcast(&event)
			}
		}
	}
}

func (sse *sse) RegisterSSE(ch chan domain.ProductChangeEvent) {
	sse.mutex.Lock()
	sse.sseClients[ch] = true
	sse.mutex.Unlock()
}

func (sse *sse) UnregisterSSE(ch chan domain.ProductChangeEvent) {
	sse.mutex.Lock()
	delete(sse.sseClients, ch)
	sse.mutex.Unlock()
}

func (sse *sse) broadcast(event *domain.ProductChangeEvent) {
	sse.mutex.RLock()
	for ch := range sse.sseClients {
		select {
		case ch <- *event:
		default:
		}
	}
	sse.mutex.RUnlock()
}

func (c *Consumer) GetSSE() *sse {
	return c.sse
}

func (c *Consumer) Close() error {
	err := c.reader.Close()
	if err != nil {
		log.Error().Err(err).Msg("Error closing kafka consumer")
		return err
	}
	log.Info().Msg("Kafka consumer closed")
	return nil
}
