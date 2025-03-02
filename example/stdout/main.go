package main

import (
	"github.com/igordth/zap-entities/stdout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	log := stdout.NewLogger()
	log.Info("some message")

	logWithCore := zap.New(
		stdout.NewCore(stdout.DefaultEncoderConfig, zapcore.DebugLevel),
	)
	logWithCore.Info("some message log with core")
}
