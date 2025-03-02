package writer

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"sync"
)

const (
	DefaultFileMode os.FileMode = 0666
	DefaultDirMode  os.FileMode = 0766
)

type File interface {
	io.Writer
	// SetCreateMode - change permission for create directories and file
	SetCreateMode(fMode, dMode os.FileMode) File
	// SetTruncateFlag - set truncate flag for os.OpenFile
	SetTruncateFlag(t bool) File
}

type file struct {
	name        string
	file        *os.File
	fCreateMode os.FileMode
	dCreateMode os.FileMode
	truncate    bool
	mu          *sync.Mutex
}

func NewFile(name string) File {
	return &file{
		name:        name,
		fCreateMode: DefaultFileMode,
		dCreateMode: DefaultDirMode,
		mu:          &sync.Mutex{},
	}
}

func (f *file) SetTruncateFlag(t bool) File {
	f.truncate = t
	return f
}

func (f *file) SetCreateMode(fMode, dMode os.FileMode) File {
	f.fCreateMode = fMode
	f.dCreateMode = dMode
	return f
}

func (f *file) getOpenFileFlag() int {
	_, err := os.Stat(f.name)
	if f.truncate || os.IsNotExist(err) {
		return os.O_TRUNC | os.O_WRONLY | os.O_CREATE
	}
	return os.O_APPEND | os.O_WRONLY
}

func (f *file) setFile() (err error) {
	// skip if already inited
	if f.file != nil {
		return
	}

	// create directories if needed
	err = os.MkdirAll(filepath.Dir(f.name), f.dCreateMode)
	if err != nil {
		return errors.Wrap(err, "os.MkdirAll")
	}

	// open file for write or create if needed
	f.file, err = os.OpenFile(f.name, f.getOpenFileFlag(), f.fCreateMode)
	if err != nil {
		return errors.Wrap(err, "os.OpenFile")
	}

	return nil
}

func (f *file) Write(p []byte) (n int, err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if err = f.setFile(); err != nil {
		return
	}
	return f.file.Write(p)
}
