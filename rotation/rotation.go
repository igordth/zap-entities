package rotation

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultTimeLayout = "2006-01-02T15:04:05.999"
)

func NewDefaultCore(file string) (zapcore.Core, zap.AtomicLevel) {
	// set rotation config/writer
	writer := &lumberjack.Logger{
		Filename:  file,
		MaxSize:   128,
		MaxAge:    90,
		LocalTime: true,
		Compress:  true,
	}
	return NewCore(writer, zapcore.InfoLevel, nil)
}

func NewCore(cfg *lumberjack.Logger, l zapcore.Level, timeEncoder zapcore.TimeEncoder) (zapcore.Core, zap.AtomicLevel) {
	// set encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(DefaultTimeLayout)
	if timeEncoder != nil {
		encoderCfg.EncodeTime = timeEncoder
	}

	// create encoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	// atomic level
	atomic := zap.NewAtomicLevelAt(l)

	// create core
	core := zapcore.NewCore(encoder, zapcore.AddSync(cfg), atomic)

	return core, atomic
}
