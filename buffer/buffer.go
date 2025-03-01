package buffer

import "go.uber.org/zap/buffer"

type Buffer interface {
	Buffer() (*buffer.Buffer, error)
}
