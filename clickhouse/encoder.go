package clickhouse

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	buf "zap-cores/buffer"
	enc "zap-cores/encoder"
)

type encoder struct {
	*EncoderConfig
	zapcore.Encoder
}

func NewEncoder(cfg EncoderConfig) zapcore.Encoder {
	return &encoder{
		&cfg,
		zapcore.NewJSONEncoder(jsonEncoderConfig),
	}
}

func (e *encoder) Clone() zapcore.Encoder {
	return &encoder{e.EncoderConfig, e.Encoder.Clone()}
}

func (e *encoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b := buf.NewSQL(e.TableName, e.LineEnding)

	// value of FieldsColumn by JSONEncoder
	fieldsEncode, err := e.Encoder.EncodeEntry(ent, fields)
	if err != nil {
		return nil, err
	}
	b.Append(e.FieldsColumn, fieldsEncode.String())

	// value of TimeColumn with encode EncodeTime
	b.Append(
		enc.NewKeyValue(e.TimeColumn).EncodeTime(ent.Time, e.EncodeTime),
	)

	// value of LevelColumn with encode EncodeLevel
	b.Append(
		enc.NewKeyValue(e.LevelColumn).EncodeLevel(ent.Level, e.EncodeLevel),
	)

	// value of NameColumn with encode EncodeName
	b.Append(
		enc.NewKeyValue(e.NameColumn).EncodeName(ent.LoggerName, e.EncodeName),
	)

	// value of MessageColumn
	b.Append(e.MessageColumn, ent.Message)

	// value of StacktraceColumn
	b.Append(e.StacktraceColumn, ent.Stack)

	if ent.Caller.Defined {
		// value of FunctionColumn
		b.Append(e.FunctionColumn, ent.Caller.Function)
		// value of CallerColumn with encode EncodeCaller
		b.Append(
			enc.NewKeyValue(e.CallerColumn).EncodeCaller(ent.Caller, e.EncodeCaller),
		)
	}

	// todo may by buf.Free()

	return b.Buffer()
}
