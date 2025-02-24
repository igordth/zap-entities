package main

import "zap-cores/stdout"

func main() {
	log := stdout.NewLogger()

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
}
