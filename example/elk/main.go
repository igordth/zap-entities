package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"zap-cores/elk"
	"zap-cores/stdout"
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
	coreElk, _ := elk.NewCore(writer, zap.InfoLevel)

	// stdout core
	coreStd, _ := stdout.NewCore(zap.DebugLevel)

	// logger with elk & stdout cores
	log := zap.New(zapcore.NewTee(coreElk, coreStd))

	log.Info("hello world")
}
