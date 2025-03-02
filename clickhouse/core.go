package clickhouse

import (
	"github.com/igordth/zap-entities/writer"
	"go.uber.org/zap/zapcore"
	"net/http"
)

func NewCore(cfg EncoderConfig, w writer.Http, level zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		NewEncoder(cfg),
		zapcore.AddSync(w),
		level,
	)
}

func NewDefaultCore(url string) zapcore.Core {
	return NewCore(
		DefaultEncoderConfig,
		writer.NewHttp(writer.HttpDefaultClient, url, http.MethodPost),
		zapcore.InfoLevel,
	)
}
