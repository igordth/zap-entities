package elk

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap/zapcore"
	"io"
)

func NewCore(cfg ecszap.EncoderConfig, w io.Writer, level zapcore.LevelEnabler) zapcore.Core {
	return ecszap.NewCore(cfg, zapcore.AddSync(w), level)
}

func NewDefaultCore(w io.Writer) zapcore.Core {
	return NewCore(ecszap.NewDefaultEncoderConfig(), w, zapcore.InfoLevel)
}
