package clog

import "io"

// Writer is an io.Writer that writes into a Handler.
// This can be used to redirect a log.Logger into a handler.
type Writer struct {
	level   LogLevel
	handler Handler
}

// NewWriter creates a new Writer to a Handler
func NewWriter(level LogLevel, handler Handler) *Writer {
	return &Writer{
		level:   level,
		handler: handler,
	}
}

// Write bytes into the logger
func (w *Writer) Write(p []byte) (n int, err error) {
	err = w.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     w.level,
		Values:    []interface{}{string(p)},
	})
	n = len(p)
	return
}

// Verify interfaces
var _ io.Writer = &Writer{}
