package clickhouse

import (
	"go.uber.org/zap/zapcore"
	"net/http"
	"zap-cores/writer"
)

func NewCore(cfg EncoderConfig, level zapcore.LevelEnabler, url string) zapcore.Core {
	return zapcore.NewCore(
		NewEncoder(cfg),
		zapcore.AddSync(writer.NewHttp(writer.HttpDefaultClient, url, http.MethodPost)),
		level,
	)
}
