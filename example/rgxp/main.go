package main

import (
	"github.com/igordth/zap-entities/file"
	"github.com/igordth/zap-entities/rgxp"
	"github.com/igordth/zap-entities/stdout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"regexp"
)

func main() {
	// basic cores
	appleCore := file.NewDefaultCore("./example/rgxp/log/apple.log")
	bananaCore := file.NewDefaultCore("./example/rgxp/log/banana.log")
	cherryCore := file.NewDefaultCore("./example/rgxp/log/cherry.log")
	stdCore := stdout.NewDefaultCore()

	// log rgxp with file cores
	log := zap.New(zapcore.NewTee(
		rgxp.NewNamedCore(appleCore, regexp.MustCompile("apple")),
		rgxp.NewNamedCore(bananaCore, regexp.MustCompile("banana")),
		rgxp.NewNamedCore(cherryCore, regexp.MustCompile("cherry")),
		rgxp.NewMessageCore(stdCore, regexp.MustCompile("stdout")),
	))

	// log to apple.log by name apple
	log.Named("apple").Info("log to apple.log by name")

	// log to apple.log & banana.log by name apple
	log.
		Named("apple").
		Named("banana").
		Info("log to apple.log & banana.log by name")

	// log to cherry.log by name and `stdout` by message
	log.Named("cherry").Info("log to cherry.log by name and `stdout` by message")
}
