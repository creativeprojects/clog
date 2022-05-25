package clog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeeHandler(t *testing.T) {
	a := NewMemoryHandler()
	b := NewMemoryHandler()
	logger := NewLogger(NewTeeHandler(a, b))
	logger.Info("test")

	assert.Equal(t, 1, len(a.log))
	assert.Equal(t, 1, len(b.log))

	assert.Equal(t, "test", a.log[0])
	assert.Equal(t, "test", b.log[0])
}

func TestTeeHandlerSetPrefix(t *testing.T) {
	a := NewMemoryHandler()
	b := NewMemoryHandler()
	logger := NewLogger(NewTeeHandler(a, b))
	logger.SetPrefix("prefix ")
	logger.Info("test")

	assert.Equal(t, 1, len(a.log))
	assert.Equal(t, 1, len(b.log))

	assert.Equal(t, "prefix test", a.log[0])
	assert.Equal(t, "prefix test", b.log[0])
}

func TestTeeHandlerReturnsError1(t *testing.T) {
	next := NewMemoryHandler()
	handler := NewTeeHandler(NewDiscardHandler(), next)
	err := handler.LogEntry(LogEntry{Values: []interface{}{"test"}})
	assert.ErrorIs(t, err, ErrMessageDiscarded)

	assert.Equal(t, 1, len(next.log))
	assert.Equal(t, "test", next.log[0])
}

func TestTeeHandlerReturnsError2(t *testing.T) {
	next := NewMemoryHandler()
	handler := NewTeeHandler(next, NewDiscardHandler())
	err := handler.LogEntry(LogEntry{Values: []interface{}{"test"}})
	assert.ErrorIs(t, err, ErrMessageDiscarded)

	assert.Equal(t, 1, len(next.log))
	assert.Equal(t, "test", next.log[0])
}

func TestTeeHandlerNil1(t *testing.T) {
	next := NewMemoryHandler()
	handler := NewTeeHandler(nil, next)
	err := handler.LogEntry(LogEntry{Values: []interface{}{"test"}})
	assert.ErrorIs(t, err, ErrNoRegisteredHandler)

	assert.Equal(t, 1, len(next.log))
	assert.Equal(t, "test", next.log[0])
}

func TestTeeHandlerNil2(t *testing.T) {
	next := NewMemoryHandler()
	handler := NewTeeHandler(next, nil)
	err := handler.LogEntry(LogEntry{Values: []interface{}{"test"}})
	assert.ErrorIs(t, err, ErrNoRegisteredHandler)

	assert.Equal(t, 1, len(next.log))
	assert.Equal(t, "test", next.log[0])
}

func ExampleTeeHandler() {
	// a tee handler is a middleware with two targets (Handler) which
	// will send all messages to both handlers as targets

	// let's create a tee handler with two TextHandler targets
	logger := NewLogger(NewTeeHandler(
		NewTextHandler("logger 1 - ", 0),
		NewTextHandler("logger 2 - ", 0),
	))
	logger.Infof("hello %s", "world")
	// Output:
	// logger 1 - hello world
	// logger 2 - hello world
}
