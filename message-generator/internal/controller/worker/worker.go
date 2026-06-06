package worker

import (
	"context"
	"log/slog"
	"time"

	"github.com/critma/message-generator/internal/usecase"
)

type workerServer struct {
	productService *usecase.ProductUsecase
	ticker         *time.Ticker
	stopQ          chan struct{}
	log            *slog.Logger
}

func NewWorkerServer(productService *usecase.ProductUsecase, tickDurration time.Duration, log *slog.Logger) *workerServer {
	ticker := time.NewTicker(tickDurration)

	return &workerServer{
		productService: productService,
		ticker:         ticker,
		stopQ:          make(chan struct{}),
		log:            log,
	}
}

func (w *workerServer) productRandomStockPrice(ctx context.Context) error {
	err := w.productService.ProductRandomStockPrice(ctx)
	if err != nil {
		return err
	}
	w.log.Info("productRandomStockPrice.message sended")
	return nil
}

func (w *workerServer) Run(ctx context.Context) error {
	for {
		select {
		case <-w.ticker.C:
			err := w.productRandomStockPrice(ctx)
			if err != nil {
				w.log.Error("productRandomStockPrice.task failed", slog.Any("error", err))
				w.Stop()
			}
		case <-w.stopQ:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *workerServer) Stop() {
	w.ticker.Stop()
	close(w.stopQ)
}
