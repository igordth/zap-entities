package main

import (
	"github.com/igordth/zap-entities/file"
	"go.uber.org/zap"
)

func main() {
	core := file.NewDefaultCore("./example/file/log/file.log")
	log := zap.New(core)
	log.Info("hello world")
}
