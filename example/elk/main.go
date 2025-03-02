package main

import (
	"github.com/igordth/zap-entities/elk"
	"github.com/igordth/zap-entities/stdout"
	"github.com/igordth/zap-entities/writer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// writer for elk core
	w := writer.NewFile("./example/elk/log/elk.log")

	// elk core
	coreElk := elk.NewDefaultCore(w)

	// stdout core
	coreStd := stdout.NewDefaultCore()

	// logger with elk & stdout cores
	log := zap.New(zapcore.NewTee(coreElk, coreStd))

	log.Info("hello world")
}
