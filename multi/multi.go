package multi

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Logger struct {
	def  *zap.Logger
	list sync.Map
}

func New(list map[string]zapcore.Core, def *zap.Logger) *Logger {
	logger := &Logger{def: def}
	for key, core := range list {
		logger.Add(key, core)
	}
	return logger
}

func (l *Logger) SetDefault(def *zap.Logger) *Logger {
	l.def = def
	return l
}

func (l *Logger) Add(key string, core zapcore.Core) *Logger {
	log := zap.NewNop().WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core {
		return core
	}))
	l.list.Store(key, log)
	return l
}

func (l *Logger) Get(key string) *zap.Logger {
	if v, ok := l.list.Load(key); ok {
		return v.(*zap.Logger)
	}
	if l.def == nil {
		return l.def
	}
	return zap.NewNop()
}
