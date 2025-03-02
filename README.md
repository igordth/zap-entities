Zap Entities
============

Cores, encoders, writers, buffers, ... for zap logger.

Basic library [go.uber.org/zap](https://github.com/uber-go/zap).

# Cores

## StdOut

Core for write logs to stdout

* `NewCore` - create core
* `NewLogger` - create logger with default options

```go
import "github.com/igordth/zap-entities/stdout"

func main() {
    log := stdout.NewLogger()
    log.Info("hello world")

    core := stdout.NewCore(stdout.DefaultEncoderConfig, zapcore.DebugLevel)
    logByCore := zap.New(core)
    logByCore.Info("hello world from core")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/examples/stdout/main.go

## ELK (Elasticsearch,Logstash,Kibana)

Creates a zapcore.Core that uses an ECS conformant JSON encoder for push logs to ELK.  
Using library [ecszap](https://pkg.go.dev/go.elastic.co/ecszap)

```go
import (
    "github.com/igordth/zap-entities/elk"
    "go.uber.org/zap"
)

func main() {
    // todo set file writer
    core := elk.NewCore(writer, zap.InfoLevel)
    log := zap.New(core)
    log.Info("hello world")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/examples/elk/main.go

## Rotation

Core for write logs to file with file rotation.  
Using library [lumberjack.v2](https://pkg.go.dev/gopkg.in/lumberjack.v2)

```go
import (
    "github.com/igordth/zap-entities/rotation"
    "go.uber.org/zap"
    "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
    coreDefault := rotation.NewDefaultCore("default.log")
    logDefault := zap.New(coreDefault)
    logDefault.Info("default log")

    coreCustom = rotation.NewCore(
        &lumberjack.Logger{
            Filename:   "custom.log",
            MaxSize:    1,
            MaxAge:     1,
            MaxBackups: 1,
            LocalTime:  true,
            Compress:   true,
        },
        rotation.DefaultEncoderConfig,
        zap.InfoLevel,
    )
    logCustom := zap.New(coreCustom)
    logCustom.Info("custom log")
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/examples/rotation/main.go

## Rgxp

Filtering logs for send to writer with `regexp.Regexp` by name or message.  
Using library [regexp](https://pkg.go.dev/regexp)

```go
import (
    "github.com/igordth/zap-entities/rgxp"
    "github.com/igordth/zap-entities/rotation"
    "github.com/igordth/zap-entities/stdout"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "regexp"
)

func main() {
    // basic cores
    appleCore := rotation.NewDefaultCore("apple.log")
    bananaCore := rotation.NewDefaultCore("banana.log")

    // rgxp cores
    rgxpLog := zap.New(zapcore.NewTee(
        rgxp.NewNamedCore(appleCore, regexp.MustCompile("apple")),
        rgxp.NewNamedCore(bananaCore, regexp.MustCompile("banana")),
    ))

    rgxpLog.Named("apple").Info("log to apple.log")
    rgxpLog.Named("banana").Info("log to banana.log")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/examples/rgxp/main.go

## Clickhouse

Core for write logs to clickhouse by [http interface](https://clickhouse.com/docs/interfaces/http).

```go
import (
    "github.com/igordth/zap-entities/clickhouse"
    "go.uber.org/zap"
)

func main() {
    core := clickhouse.NewCore(clickhouse.DefaultEncoderConfig, zap.InfoLevel, "http://localhost:8123")
    log := zap.New(core)
    log.Info("hello world")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/examples/clickhouse/main.go

# Encoder package

Contain entities that are used in encoder.

## Field

Field encoder.  
Encode field of `zapcore.Entry` by functions:

* `zapcore.TimeEncoder` - for field Time
* `zapcore.LevelEncoder` - for field Level
* `zapcore.CallerEncoder` - for field Caller
* `zapcore.NameEncoder` - for field LoggerName

Methods:

* `Time(inp time.Time, fn zapcore.TimeEncoder) string`
* `Level(inp zapcore.Level, fn zapcore.LevelEncoder) string`
* `Caller(inp zapcore.EntryCaller, fn zapcore.CallerEncoder) string`
* `Name(inp string, fn zapcore.NameEncoder) string`

Options:

* `Value` - current value


It is mainly used in `zapcore.Encored` with `buffer.SQL`

```go
import (
   "github.com/igordth/zap-entities/buffer"
   "github.com/igordth/zap-entities/encoder"
   "go.uber.org/zap/zapcore"
)

// ...

func (e *myEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
    // init buffer
    tableName := "logs"
    lineEnding := ";\n"
    b := bufffer.NewSQL(tableName, lineEnding)

    // init encoder.Field & encode
    fieldEncoder := new(encoder.Field)
    encodedTime := fieldEncoder.Time(ent.Time, e.EncodeTime)
	
    // append encoded field to buffer
    b.Append(e.TimeColumn, encodedTime))
}
```

## EmptyPrimitiveArray

Array encoder implemented of 
[zapcore.PrimitiveArrayEncoder](https://pkg.go.dev/go.uber.org/zap/zapcore#PrimitiveArrayEncoder) used in 
encoderFunctions like `zapcore.TimeEncoder`... with empty methods.

## ValuePrimitiveArray

Array encoder implemented of 
[zapcore.PrimitiveArrayEncoder](https://pkg.go.dev/go.uber.org/zap/zapcore#PrimitiveArrayEncoder) used in
encoderFunctions like `zapcore.TimeEncoder`... set value to option `Value` in AppendFunc.

Currently used in `encoder.Field` to save encoded value.

Options:

* `Value` - current value

# Buffer

Buffer package. Used in `zapcore.Encoder` (add to buffer) and `zapcore.WriteSyncer` (write buffer to output).

## SQL

Buffer for sql queries.

Methods:

* `Append(column string, value string)` - add column and it`s value to buffer
* `Buffer() (*buffer.Buffer, error)` - buffered result sql query

# Writer

Writers for `zapcore.Core` which belong to the interface `zapcore.WriteSyncer` or `io.Writer`.  
Used in `zapcore.NewCore(encoder, WRITER, lebelEncoder)`

## Http

Writer http - send logs by http.

```go
import (
    "github.com/igordth/zap-entities/writer"
    "go.uber.org/zap/zapcore"
    "net/http"
)

func main() {
    // init writer
    httpWriter := writer.NewHttp(writer.HttpDefaultClient, "http://localhost:8123", http.MethodPost)

    // init core with writer httpWriter
    core := zapcore.NewCore(
        clickhouse.NewEncoder(clickhouse.DefaultEncoderConfig),
        zapcore.AddSync(httpWriter),
        zap.InfoLevel,
    )
	
    // init log
    log := zap.New(core)
    log.Info("hello world")
}
```

# Sources:

* https://github.com/uber-go/zap
* https://pkg.go.dev/go.elastic.co/ecszap
* https://pkg.go.dev/gopkg.in/lumberjack.v2
* https://pkg.go.dev/regexp
* https://clickhouse.com/docs/interfaces/http

# TODO

1. Broker
   1. RabbitMQ
   2. Kafka
2. Telegram
3. Grpc writer
4. Simple file writer
5. check creating dir in rotation on alpine