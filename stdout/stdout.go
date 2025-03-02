package stdout

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	DefaultEncoderConfig = zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "M",
		StacktraceKey: "S",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.999"),
	}
)

func NewCore(encoderCfg zapcore.EncoderConfig, level zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		level,
	)
}

func NewDefaultCore() zapcore.Core {
	return NewCore(DefaultEncoderConfig, zapcore.DebugLevel)
}

func NewLogger() *zap.Logger {
	return zap.New(NewDefaultCore())
}
