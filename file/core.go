package file

import (
	"github.com/igordth/zap-entities/writer"
	"go.uber.org/zap/zapcore"
)

var (
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

func NewCore(encCfg zapcore.EncoderConfig, w writer.File, level zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encCfg),
		zapcore.AddSync(w),
		level,
	)
}

func NewDefaultCore(fileName string) zapcore.Core {
	return NewCore(
		DefaultEncoderConfig,
		writer.NewFile(fileName),
		zapcore.InfoLevel,
	)
}
