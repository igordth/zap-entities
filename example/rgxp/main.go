package main

import (
	"github.com/igordth/zap-entities/rgxp"
	"github.com/igordth/zap-entities/rotation"
	"github.com/igordth/zap-entities/stdout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"regexp"
)

func main() {
	// basic cores
	appleCore := rotation.NewDefaultCore("./example/rgxp/log/apple.log")
	bananaCore := rotation.NewDefaultCore("./example/rgxp/log/banana.log")
	cherryCore := rotation.NewDefaultCore("./example/rgxp/log/cherry.log")
	stdCore := stdout.NewCore(stdout.DefaultEncoderConfig, zap.InfoLevel)

	// rgxp cores
	rgxpLog := zap.New(zapcore.NewTee(
		rgxp.NewNamedCore(appleCore, regexp.MustCompile("apple")),
		rgxp.NewNamedCore(bananaCore, regexp.MustCompile("banana")),
		rgxp.NewNamedCore(cherryCore, regexp.MustCompile("cherry")),
		rgxp.NewMessageCore(stdCore, regexp.MustCompile("stdout")),
	))

	// log to apple.log by name apple
	rgxpLog.Named("apple").Info("log to apple.log by name")

	// log to apple.log & banana.log by name apple
	rgxpLog.
		Named("apple").
		Named("banana").
		Info("log to apple.log & banana.log by name")

	// log to cherry.log by name and `stdout` by message
	rgxpLog.Named("cherry").Info("log to cherry.log by name and `stdout` by message")
}
