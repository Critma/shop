package zlog

import "github.com/rs/zerolog/log"

func PrintErr(method string, input any, err error, message string) {
	log.Error().Str("method", method).Any("input", input).Err(err).Msg(message)
}

func PrintInfo(method string, data any, message string) {
	log.Info().Str("method", method).Any("data", data).Msg(message)
}
