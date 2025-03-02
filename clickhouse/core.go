package clickhouse

import (
	"github.com/igordth/zap-entities/writer"
	"go.uber.org/zap/zapcore"
	"net/http"
)

func NewCore(cfg EncoderConfig, level zapcore.LevelEnabler, url string) zapcore.Core {
	return zapcore.NewCore(
		NewEncoder(cfg),
		zapcore.AddSync(writer.NewHttp(writer.HttpDefaultClient, url, http.MethodPost)),
		level,
	)
}
