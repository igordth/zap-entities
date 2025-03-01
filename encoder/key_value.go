package encoder

import (
	"go.uber.org/zap/zapcore"
	"time"
	"zap-cores/convtostr"
)

type KeyValue struct {
	Key string
	ValuePrimitiveArray
}

func NewKeyValue(key string) *KeyValue {
	return &KeyValue{Key: key}
}

func (c *KeyValue) EncodeTime(inp time.Time, fn zapcore.TimeEncoder) (string, string) {
	switch {
	case fn == nil:
		c.Value = convtostr.AnyStringer(inp)
	default:
		fn(inp, c)
	}
	return c.Key, c.Value
}

func (c *KeyValue) EncodeLevel(inp zapcore.Level, fn zapcore.LevelEncoder) (string, string) {
	switch {
	case fn == nil:
		c.Value = convtostr.AnyStringer(inp)
	default:
		fn(inp, c)
	}
	return c.Key, c.Value
}

func (c *KeyValue) EncodeCaller(inp zapcore.EntryCaller, fn zapcore.CallerEncoder) (string, string) {
	switch {
	case fn == nil:
		c.Value = convtostr.AnyStringer(inp)
	default:
		fn(inp, c)
	}
	return c.Key, c.Value
}

func (c *KeyValue) EncodeName(inp string, fn zapcore.NameEncoder) (string, string) {
	switch {
	case fn == nil:
		c.Value = convtostr.AnyStringer(inp)
	default:
		fn(inp, c)
	}
	return c.Key, c.Value
}
