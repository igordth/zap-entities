package rotation

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	DefaultWriter = func(file string) *lumberjack.Logger {
		return &lumberjack.Logger{
			Filename:  file,
			MaxSize:   128,
			MaxAge:    90,
			LocalTime: true,
			Compress:  true,
		}
	}
	DefaultEncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.999"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

func NewDefaultCore(file string) zapcore.Core {
	return NewCore(
		DefaultWriter(file),
		DefaultEncoderConfig,
		zapcore.InfoLevel,
	)
}

func NewCore(w *lumberjack.Logger, encCfg zapcore.EncoderConfig, level zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encCfg),
		zapcore.AddSync(w),
		level,
	)
}
