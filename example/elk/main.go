package main

import (
	"github.com/igordth/zap-entities/elk"
	"github.com/igordth/zap-entities/stdout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// writer for elk core
	writer := &lumberjack.Logger{
		Filename:   "./example/elk/log/elk.log",
		MaxSize:    128,
		MaxAge:     7,
		MaxBackups: 1,
		LocalTime:  true,
		Compress:   false,
	}

	// elk core
	coreElk := elk.NewCore(writer, zap.InfoLevel)

	// stdout core
	coreStd := stdout.NewCore(stdout.DefaultEncoderConfig, zap.DebugLevel)

	// logger with elk & stdout cores
	log := zap.New(zapcore.NewTee(coreElk, coreStd))

	log.Info("hello world")
}
