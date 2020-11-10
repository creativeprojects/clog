package clog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorShouldSendToBackup(t *testing.T) {
	backup := NewMemoryHandler()
	handler := NewSafeHandler(NewDiscardHandler(), backup)
	err := handler.LogEntry(LogEntry{})
	assert.NoError(t, err)
	assert.Len(t, backup.log, 1)
}

func TestDoubleErrorShouldReturnError(t *testing.T) {
	handler := NewSafeHandler(NewDiscardHandler(), NewDiscardHandler())
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}

func TestSafeHandlerShouldFailPrimaryHandler(t *testing.T) {
	handler := NewSafeHandler(nil, NewDiscardHandler())
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}

func TestSafeHandlerShouldFailSecondaryHandler(t *testing.T) {
	handler := NewSafeHandler(NewDiscardHandler(), nil)
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}

func TestSafeHandlerCanCanSetPrefix(t *testing.T) {
	memHandler := NewMemoryHandler()
	filter := NewSafeHandler(memHandler, memHandler)
	filter.SetPrefix("_test_")
	filter.LogEntry(NewLogEntry(3, LevelInfo, "hello world"))
	assert.Equal(t, "_test_hello world", memHandler.log[0])
}

func ExampleSafeHandler() {
	// a safe handler is a middleware with two targets (Handler):
	// - it will send all messages to the primary handler
	// - it will only send the messages again to the backup handler if the primary returned an error

	// let's create a safe handler with these two targets: a DiscardHandler and a TextHandler
	// the DiscardHandler is a special handler that always discards your log messages and returns an error
	logger := NewLogger(NewSafeHandler(NewDiscardHandler(), NewTextHandler("backup ", 0)))
	logger.Infof("hello %s", "world")
	// Output: backup hello world
}
