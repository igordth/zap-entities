package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"regexp"
	"zap-cores/rgxp"
	"zap-cores/rotation"
	"zap-cores/stdout"
)

func main() {
	// define cores
	appleCore, _ := rotation.NewDefaultCore("./example/rgxp/log/apple.log")
	bananaCore, _ := rotation.NewDefaultCore("./example/rgxp/log/banana.log")
	cherryCore, _ := rotation.NewDefaultCore("./example/rgxp/log/cherry.log")
	stdCore, _ := stdout.NewCore(zap.InfoLevel)

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
