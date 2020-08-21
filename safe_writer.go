package clog

import (
	"io"
	"sync"
)

// SafeWriter is a thread safe io.Writer
type SafeWriter struct {
	writer io.Writer
	mu     sync.Mutex
}

// NewSafeWriter creates a thread-safe io.Writer
func NewSafeWriter(writer io.Writer) *SafeWriter {
	return &SafeWriter{
		writer: writer,
		mu:     sync.Mutex{},
	}
}

// Write data (thread safe)
func (w *SafeWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.Write(p)
}

// Verify interface
var _ io.Writer = &SafeWriter{}
