package clog

import (
	"sync"
)

const defaultCapacity = 100

// AsyncHandler asynchronously send log messages to the next handler in the chain
type AsyncHandler struct {
	*middlewareHandler
	wg      sync.WaitGroup
	entries chan LogEntry
	done    chan interface{}
	closed  bool
}

// NewAsyncHandler returns a handler that sends logs asynchronously.
func NewAsyncHandler(destination Handler) *AsyncHandler {
	return NewAsyncHandlerWithCapacity(destination, defaultCapacity)
}

// NewAsyncHandlerWithCapacity returns a handler that sends logs asynchronously.
func NewAsyncHandlerWithCapacity(next Handler, capacity uint) *AsyncHandler {
	entries := make(chan LogEntry, capacity)
	done := make(chan interface{})
	handler := &AsyncHandler{
		middlewareHandler: newMiddlewareHandler(next),
		wg:                sync.WaitGroup{},
		entries:           entries,
		done:              done,
		closed:            false,
	}
	// start the goroutine to handle messages in the background
	go func(handler Handler, entries chan LogEntry, done chan interface{}) {
		for logEntry := range entries {
			_ = handler.LogEntry(logEntry)
		}
		// the entries channels has been drained
		close(done)
	}(next, entries, done)
	return handler
}

// LogEntry sends the log message asynchronously to the next handler.
// Please note the call will block when the buffer capacity is reached.
// It will return an error if there's no "next" handler, of if we called the Close method.
// Otherwise it will always return nil as it doesn't know if the message will be delivered.
func (h *AsyncHandler) LogEntry(logEntry LogEntry) error {
	// make sure we don't keep sending messages to a closed channel
	if h.closed {
		return ErrHandlerClosed
	}
	if h.next == nil {
		return ErrNoRegisteredHandler
	}
	h.entries <- logEntry
	return nil
}

// SetPrefix sets a prefix on every log message
func (h *AsyncHandler) SetPrefix(prefix string) Handler {
	if h.next == nil {
		return h
	}
	h.next.SetPrefix(prefix)
	return h
}

// Close blocks until all log messages have been delivered
func (h *AsyncHandler) Close() {
	h.closed = true
	close(h.entries)
	<-h.done
}

// Verify interface
var (
	_ Handler           = &AsyncHandler{}
	_ MiddlewareHandler = &AsyncHandler{}
)
