package clog

import (
	"io"
	"sync"
)

// syncWriter is a thread safe io.Writer
type syncWriter struct {
	io.Writer
	sync.Mutex
}

// Write data (thread safe)
func (w *syncWriter) Write(p []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	return w.Writer.Write(p)
}

// fdWriter is an io.Writer that also has an Fd method. The most common
// example of an fdWriter is an *os.File.
type fdWriter interface {
	io.Writer
	Fd() uintptr
}

// syncFdWriter is a thread safe io.Writer with Fd()
// Inspired by https://github.com/go-kit/log/
type syncFdWriter struct {
	fdWriter
	sync.Mutex
}

// Write data (thread safe)
func (w *syncFdWriter) Write(p []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	return w.fdWriter.Write(p)
}

// NewSyncWriter creates a thread-safe io.Writer
func NewSyncWriter(writer io.Writer) io.Writer {
	fdWriter, ok := writer.(fdWriter)
	if ok {
		return &syncFdWriter{
			fdWriter: fdWriter,
		}
	}
	return &syncWriter{
		Writer: writer,
	}
}

// Verify interface
var (
	_ io.Writer = &syncWriter{}
	_ io.Writer = &syncFdWriter{}
)
