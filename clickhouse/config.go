package clickhouse

import (
	"go.uber.org/zap/zapcore"
)

type EncoderConfig struct {
	TableName          string                `json:"tableName" yaml:"tableName"`
	TimeColumn         string                `json:"timeColumn" yaml:"timeColumn"`
	LevelColumn        string                `json:"levelColumn" yaml:"levelColumn"`
	NameColumn         string                `json:"nameColumn" yaml:"nameColumn"`
	CallerColumn       string                `json:"callerColumn" yaml:"callerColumn"`
	FunctionColumn     string                `json:"functionColumn" yaml:"functionColumn"`
	MessageColumn      string                `json:"messageColumn" yaml:"messageColumn"`
	StacktraceColumn   string                `json:"stacktraceColumn" yaml:"stacktraceColumn"`
	FieldsColumn       string                `json:"fieldsColumn" yaml:"fieldsColumn"`
	LineEnding         string                `json:"lineEnding" yaml:"lineEnding"`
	EncodeLevel        zapcore.LevelEncoder  `json:"levelEncoder" yaml:"levelEncoder"`
	EncodeTime         zapcore.TimeEncoder   `json:"timeEncoder" yaml:"timeEncoder"`
	EncodeCaller       zapcore.CallerEncoder `json:"callerEncoder" yaml:"callerEncoder"`
	EncodeName         zapcore.NameEncoder   `json:"nameEncoder" yaml:"nameEncoder"`
	JsonEncodeTime     zapcore.TimeEncoder
	JsonEncodeDuration zapcore.DurationEncoder
}

var (
	DefaultEncoderConfig = EncoderConfig{
		TableName:        "logs",
		TimeColumn:       "date_time",
		LevelColumn:      "level",
		NameColumn:       "name",
		CallerColumn:     "caller",
		FunctionColumn:   "function",
		MessageColumn:    "message",
		StacktraceColumn: "stack",
		FieldsColumn:     "fields",
		LineEnding:       ";\n",
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
		EncodeCaller:     zapcore.ShortCallerEncoder,
		JsonEncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.999"),
	}

	jsonEncoderConfig = zapcore.EncoderConfig{
		TimeKey:        zapcore.OmitKey,
		LevelKey:       zapcore.OmitKey,
		NameKey:        zapcore.OmitKey,
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     zapcore.OmitKey,
		StacktraceKey:  zapcore.OmitKey,
		SkipLineEnding: true,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)
