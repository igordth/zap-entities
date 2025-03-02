package main

import (
	"errors"
	"github.com/igordth/zap-entities/clickhouse"
	"go.uber.org/zap"
)

func main() {
	log := zap.New(
		clickhouse.NewCore(clickhouse.DefaultEncoderConfig, zap.InfoLevel, "http://localhost:8123"),
	)

	log.
		Named("name").
		With(
			zap.Int("int", 1),
			zap.String("string", "some string"),
			zap.Bool("bool", true),
		).
		Info("hello world", zap.Bool("bool2", true))

	log.Error("message 'to escape'", zap.Error(errors.New("some error")))
}
