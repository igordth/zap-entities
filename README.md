Zap Entities
============

Cores, encoders, writers, buffers, ... for zap logger.

Basic library [go.uber.org/zap](https://github.com/uber-go/zap).

# Cores

## StdOut

Core for write logs to stdout

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
_Example in:_ https://github.com/igordth/zap-entities/tree/master/example/stdout/main.go

## File

Core for write logs to file.

```go
import (
    "github.com/igordth/zap-entities/file"
    "go.uber.org/zap"
)

func main() {
    core := file.NewDefaultCore("file.log")
    log := zap.New(core)
    log.Info("hello world")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/example/file/main.go

## ELK (Elasticsearch,Logstash,Kibana)

Creates a zapcore.Core that uses an ECS conformant JSON encoder for push logs to ELK.  
Using library [ecszap](https://pkg.go.dev/go.elastic.co/ecszap)

```go
import (
    "github.com/igordth/zap-entities/elk"
    "github.com/igordth/zap-entities/writer"
    "go.uber.org/zap"
)

func main() {
    w := writer.NewFile("./example/elk/log/elk.log")
    core := elk.NewDefaultCore(writer, zap.InfoLevel)
    log := zap.New(core)
    log.Info("hello world")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/example/elk/main.go

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
    core := rotation.NewDefaultCore("default.log")
    log := zap.New(coreDefault)
    log.Info("default log")
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/example/rotation/main.go

## Rgxp

Filtering logs for send to writer with `regexp.Regexp` by name or message.  
Using library [regexp](https://pkg.go.dev/regexp)

```go
import (
    "github.com/igordth/zap-entities/file"
    "github.com/igordth/zap-entities/rgxp"
    "github.com/igordth/zap-entities/stdout"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "regexp"
)

func main() {
    // basic cores
    appleCore := file.NewDefaultCore("apple.log")
    bananaCore := file.NewDefaultCore("banana.log")

    // log rgxp with file cores
    log := zap.New(zapcore.NewTee(
        rgxp.NewNamedCore(appleCore, regexp.MustCompile("apple")),
        rgxp.NewNamedCore(bananaCore, regexp.MustCompile("banana")),
    ))

    log.Named("apple").Info("log to apple.log")
    log.Named("banana").Info("log to banana.log")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/example/rgxp/main.go

## Clickhouse

Core for write logs to clickhouse by [http interface](https://clickhouse.com/docs/interfaces/http).

```go
import (
    "github.com/igordth/zap-entities/clickhouse"
    "go.uber.org/zap"
)

func main() {
    core := clickhouse.NewDefaultCore("http://localhost:8123")
    log := zap.New(core)
    log.Info("hello world")
}
```

_Example in:_ https://github.com/igordth/zap-entities/tree/master/example/clickhouse/main.go

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

## File

Writer file - save logs to file.

Interface

```go
type File interface {
    io.Writer
    // SetCreateMode - change permission for create directories and file
    SetCreateMode(fMode, dMode os.FileMode) File
    // SetTruncateFlag - set truncate flag for os.OpenFile
    SetTruncateFlag(t bool) File
}
```

Example

```go
import (
    "github.com/igordth/zap-entities/writer"
)

func main() {
    w := writer.NewFile("./example/file/log/file.log")
    // todo ...
}
```

## Http

Writer http - send logs by http.

```go
import (
    "github.com/igordth/zap-entities/writer"
    "net/http"
)

func main() {
    w := writer.NewHttp(writer.HttpDefaultClient, "http://localhost:8123", http.MethodPost)
    // todo ...
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
4. Limited by size file writer
5. check creating dir in rotation on alpine
6. lumberjack.v2 use fork with normal file mode