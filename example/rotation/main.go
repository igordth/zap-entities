package main

import (
	"github.com/igordth/zap-entities/rotation"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// default core
	core := rotation.NewDefaultCore("./example/rotation/log/default.log")
	log := zap.New(core)
	for i := 0; i < 100; i++ {
		log.Info("loop log", zap.Int("i", i))
	}

	// core with self config
	atomic := zap.NewAtomicLevelAt(zap.InfoLevel)
	fishString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt..."
	core = rotation.NewCore(
		&lumberjack.Logger{
			Filename:   "./example/rotation/log/custom.log",
			MaxSize:    1,
			MaxAge:     1,
			MaxBackups: 1,
			LocalTime:  true,
			Compress:   true,
		},
		rotation.DefaultEncoderConfig,
		atomic,
	)
	log = zap.New(core)
	for i := 1; i <= 7000; i++ {
		log := log.With(zap.Int("i", i), zap.String("fishString", fishString))
		log.Info("loop log")
		log.Debug("loop log") // write only after 50 line
		if i == 50 {
			atomic.SetLevel(zap.DebugLevel) // turn on debug level
		}
	}
}
