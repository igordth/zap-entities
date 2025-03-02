package elk

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap/zapcore"
	"io"
)

func NewCore(w io.Writer, level zapcore.LevelEnabler) zapcore.Core {
	return ecszap.NewCore(
		ecszap.NewDefaultEncoderConfig(),
		zapcore.AddSync(w),
		level,
	)
}
