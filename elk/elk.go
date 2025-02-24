package elk

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func NewCore(w io.Writer, l zapcore.Level) (c zapcore.Core, a zap.AtomicLevel) {
	a = zap.NewAtomicLevelAt(l)
	c = ecszap.NewCore(
		ecszap.NewDefaultEncoderConfig(),
		zapcore.AddSync(w),
		a,
	)
	return
}
