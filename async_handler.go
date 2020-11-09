package clog

import (
	"errors"
	"sync"
)

// AsyncHandler forgets any log message
type AsyncHandler struct {
	handler Handler
	wg      sync.WaitGroup
}

// NewAsyncHandler returns a handler that sends logs asynchronously.
func NewAsyncHandler(destination Handler) *AsyncHandler {
	return &AsyncHandler{
		handler: destination,
		wg:      sync.WaitGroup{},
	}
}

// SetHandler sets a new handler for the filter
func (l *AsyncHandler) SetHandler(handler Handler) {
	l.handler = handler
}

// GetHandler returns the current handler used by the filter
func (l *AsyncHandler) GetHandler() Handler {
	return l.handler
}

// LogEntry sends the log message asynchronously to the next handler.
// It will return an error if there's no "next" handler
// Otherwise it will always return nil as it doesn't know if the message will be delivered.
func (l *AsyncHandler) LogEntry(logEntry LogEntry) error {
	if l.handler == nil {
		return errors.New("no registered handler")
	}
	l.wg.Add(1)
	go func(logEntry LogEntry) {
		l.handler.LogEntry(logEntry)
		l.wg.Done()
	}(logEntry)
	return nil
}

// SetPrefix sets a prefix on every log message
func (l *AsyncHandler) SetPrefix(prefix string) {
	if l.handler == nil {
		return
	}
	l.handler.SetPrefix(prefix)
}

// Wait blocks until all log messages have been delivered
func (l *AsyncHandler) Wait() {
	l.wg.Wait()
}

// Verify interface
var (
	_ Handler      = &AsyncHandler{}
	_ HandlerChain = &AsyncHandler{}
)
