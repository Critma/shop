package zlog

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		log.WithLevel(zerolog.Level(lvl)).Fields(fields).Msg(msg)
	})
}
