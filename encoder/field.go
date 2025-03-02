package encoder

import (
	"github.com/igordth/tostring"
	"go.uber.org/zap/zapcore"
	"time"
)

type Field struct {
	ValuePrimitiveArray
}

func (c *Field) Time(inp time.Time, fn zapcore.TimeEncoder) string {
	switch {
	case fn == nil:
		c.Value = tostring.Stringer(inp)
	default:
		fn(inp, c)
	}
	return c.Value
}

func (c *Field) Level(inp zapcore.Level, fn zapcore.LevelEncoder) string {
	switch {
	case fn == nil:
		c.Value = tostring.Stringer(inp)
	default:
		fn(inp, c)
	}
	return c.Value
}

func (c *Field) Caller(inp zapcore.EntryCaller, fn zapcore.CallerEncoder) string {
	switch {
	case fn == nil:
		c.Value = tostring.Stringer(inp)
	default:
		fn(inp, c)
	}
	return c.Value
}

func (c *Field) Name(inp string, fn zapcore.NameEncoder) string {
	switch {
	case fn == nil:
		c.Value = tostring.Stringer(inp)
	default:
		fn(inp, c)
	}
	return c.Value
}
