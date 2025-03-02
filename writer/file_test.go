package writer

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestFile_Write(t *testing.T) {
	writer := NewFile("log/file.log")
	content := []byte("hello world\n")
	n, err := writer.Write(content)
	assert.NilError(t, err)
	assert.Equal(t, n, len(content))
}
