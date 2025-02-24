package main

import (
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"zap-cores/rotation"
)

func main() {
	// default core
	file := "./example/rotation/log/default.log"
	core, atomic := rotation.NewDefaultCore(file)
	log := zap.New(core)
	for i := 0; i <= 100; i++ {
		log.Info("loop log", zap.Int("i", i))
		log.Debug("loop log", zap.Int("i", i)) // not write before i <= 50, by default InfoLevel
		if i == 50 {
			// change log level to debug
			atomic.SetLevel(zap.DebugLevel)
		}
	}

	// core with self config
	fishString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	core, _ = rotation.NewCore(
		&lumberjack.Logger{
			Filename:   "./example/rotation/log/rotation.log",
			MaxSize:    1,
			MaxAge:     1,
			MaxBackups: 1,
			LocalTime:  true,
			Compress:   true,
		},
		zap.InfoLevel,
		nil,
	)
	log = zap.New(core)
	for i := 0; i <= 20000; i++ {
		log.Info("loop log", zap.Int("i", i), zap.String("fishString", fishString))
	}
}
