package product_update_subscribe

import (
	"context"
	"shopapi/internal/adapter/kafka_consume"
	"shopapi/internal/domain"
	"time"

	"github.com/rs/zerolog/log"
)

type Usecase struct {
	consumer *kafka_consume.Consumer
}

var usecase *Usecase

func NewUsecase(comsumer *kafka_consume.Consumer) *Usecase {
	uc := &Usecase{
		consumer: comsumer,
	}
	usecase = uc
	return uc
}

// func (u *Usecase) ProductUpdateSubscribeWs(ctx context.Context, ws *websocket.Conn) error {
// 	u.socket.AddConnection(ws)
// 	return nil
// }

func (u *Usecase) ListenSSEBroadcast(ctx context.Context) chan domain.ProductChangeEvent {
	const timeout = 30 * time.Second
	productChan := make(chan domain.ProductChangeEvent)
	outputChan := make(chan domain.ProductChangeEvent)
	u.consumer.GetSSE().RegisterSSE(productChan)
	log.Info().Msg("Listening for SSE broadcast")

	go func() {
		for {
			select {
			case <-ctx.Done():
				u.closeSEE(productChan, outputChan)
				return
			case <-time.After(timeout):
				u.closeSEE(productChan, outputChan)
				return
			case productEvent := <-productChan:
				outputChan <- productEvent
			}
		}
	}()
	return outputChan
}

func (u *Usecase) closeSEE(productChan, outputChan chan domain.ProductChangeEvent) {
	u.consumer.GetSSE().UnregisterSSE(productChan)
	close(outputChan)
	log.Info().Msg("Stopped listening for SSE broadcast")
}
