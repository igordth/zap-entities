package main

import (
	"go.uber.org/zap"
	"zap-cores/multi"
	"zap-cores/rotation"
)

func main() {
	/* ---BASIC USAGE--- */

	// define cores
	appleCore, _ := rotation.NewDefaultCore("./example/multi/log/apple.log")
	bananaCore, _ := rotation.NewDefaultCore("./example/multi/log/banana.log")
	cherryCore, _ := rotation.NewDefaultCore("./example/multi/log/cherry.log")

	// init loggers
	appleLog := zap.New(appleCore)
	bananaLog := zap.New(bananaCore)
	cherryLog := zap.New(cherryCore)

	// logs
	appleLog.Info("log to apple.log")
	bananaLog.Info("log to banana.log")
	cherryLog.Info("log to cherry.log")

	/* ---USAGE WITH PACKAGE MULTI--- */

	multiLog := new(multi.Logger).
		Add("apple", appleCore).
		Add("banana", bananaCore).
		Add("cherry", cherryCore)

	multiLog.Get("apple").Info("log to apple.log from multi")
	multiLog.Get("banana").Info("log to banana.log from multi")
	multiLog.Get("cherry").Info("log to cherry.log from multi")
}
