package clog

import (
	"errors"
	"sync"
)

const defaultCapacity = 100

// AsyncHandler forgets any log message
type AsyncHandler struct {
	handler Handler
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
func NewAsyncHandlerWithCapacity(destination Handler, capacity uint) *AsyncHandler {
	entries := make(chan LogEntry, capacity)
	done := make(chan interface{})
	handler := &AsyncHandler{
		handler: destination,
		wg:      sync.WaitGroup{},
		entries: entries,
		done:    done,
		closed:  false,
	}
	// start the goroutine to handle messages in the background
	go func(handler Handler, entries chan LogEntry, done chan interface{}) {
		for logEntry := range entries {
			handler.LogEntry(logEntry)
		}
		// the entries channels has been drained
		close(done)
	}(destination, entries, done)
	return handler
}

// SetHandler sets a new handler for the filter
func (h *AsyncHandler) SetHandler(handler Handler) {
	h.handler = handler
}

// GetHandler returns the current handler used by the filter
func (h *AsyncHandler) GetHandler() Handler {
	return h.handler
}

// LogEntry sends the log message asynchronously to the next handler.
// Please note the call will block when the buffer capacity is reached.
// It will return an error if there's no "next" handler, of if we called the Close method.
// Otherwise it will always return nil as it doesn't know if the message will be delivered.
func (h *AsyncHandler) LogEntry(logEntry LogEntry) error {
	// make sure we don't keep sending messages to a closed channel
	if h.closed {
		return errors.New("handler is closed")
	}
	if h.handler == nil {
		return errors.New("no registered handler")
	}
	h.entries <- logEntry
	return nil
}

// SetPrefix sets a prefix on every log message
func (h *AsyncHandler) SetPrefix(prefix string) {
	if h.handler == nil {
		return
	}
	h.handler.SetPrefix(prefix)
}

// Close blocks until all log messages have been delivered
func (h *AsyncHandler) Close() {
	h.closed = true
	close(h.entries)
	<-h.done
}

// Verify interface
var (
	_ Handler      = &AsyncHandler{}
	_ HandlerChain = &AsyncHandler{}
)
