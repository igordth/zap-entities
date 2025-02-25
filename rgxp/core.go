package rgxp

import (
	"go.uber.org/zap/zapcore"
	"regexp"
)

func NewNamedCore(c zapcore.Core, r *regexp.Regexp) zapcore.Core {
	return &nameCore{c, r}
}

type nameCore struct {
	zapcore.Core
	regexp *regexp.Regexp
}

// Check determines whether the supplied Entry should be logged
func (c *nameCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) && c.regexp.MatchString(ent.LoggerName) {
		return ce.AddCore(ent, c)
	}
	return ce
}

type messageCore struct {
	zapcore.Core
	regexp *regexp.Regexp
}

func NewMessageCore(c zapcore.Core, r *regexp.Regexp) zapcore.Core {
	return &messageCore{c, r}
}

// Check determines whether the supplied Entry should be logged
func (c *messageCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) && c.regexp.MatchString(ent.Message) {
		return ce.AddCore(ent, c)
	}
	return ce
}
