package buffer

import (
	"errors"
	"fmt"
	"go.uber.org/zap/buffer"
	"strings"
)

var (
	ErrSqlTableNameEmpty = errors.New("table name is empty")
	ErrSqlValOrColEmpty  = errors.New("empty values/columns")
	ErrSqlValOrColLenMis = errors.New("values/columns length miss")
)

type SQL interface {
	Buffer
	Append(key string, value string)
}

type sql struct {
	buf        *buffer.Buffer
	tableName  string
	lineEnding string
	columns    []string
	values     []string
}

func NewSQL(tableName string, lineEnding string) SQL {
	return &sql{
		buf:        buffer.NewPool().Get(),
		tableName:  tableName,
		lineEnding: lineEnding,
	}
}

func (q *sql) Append(key string, value string) {
	if key == "" || value == "" {
		return
	}
	value = strings.Replace(value, "'", "\\'", -1)
	q.columns = append(q.columns, key)
	q.values = append(q.values, value)
}

func (q *sql) Buffer() (*buffer.Buffer, error) {
	switch {
	case q.tableName == "":
		return q.buf, ErrSqlTableNameEmpty
	case len(q.values) != len(q.columns):
		return q.buf, ErrSqlValOrColLenMis
	case len(q.values) == 0:
		return q.buf, ErrSqlValOrColEmpty
	}
	query := fmt.Sprintf(
		`INSERT INTO %s(%s)VALUES('%s')`,
		q.tableName,
		strings.Join(q.columns, ","),
		strings.Join(q.values, "','"),
	)
	q.buf.AppendString(query)
	q.buf.AppendString(q.lineEnding)
	return q.buf, nil
}
