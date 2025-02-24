package stdout

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	DefaultTimeLayout = "2006-01-02T15:04:05.999"
)

func NewCore(l zapcore.Level) (zapcore.Core, zap.AtomicLevel) {
	// set encoder config
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(DefaultTimeLayout)
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// create encoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	// atomic level
	atomic := zap.NewAtomicLevelAt(l)

	// create core
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), atomic)

	return core, atomic
}

func NewLogger() *zap.Logger {
	core, _ := NewCore(zapcore.DebugLevel)
	return zap.New(core)
}
