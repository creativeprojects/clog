package clog

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleHandler(t *testing.T) {
	handler := NewConsoleHandler("", 0)
	for level := LevelDebug; level <= LevelError; level++ {
		err := handler.LogEntry(LogEntry{
			Level:  level,
			Values: []interface{}{level},
		})
		assert.NoError(t, err)
	}
}

func TestConsoleHandlerPrefix(t *testing.T) {
	buffer := &strings.Builder{}
	handler := NewConsoleHandler("", 0)
	// manually change the output to our local buffer
	handler.logger.SetOutput(buffer)

	err := handler.LogEntry(NewLogEntry(0, LevelInfo, "hello one"))
	assert.NoError(t, err)
	handler.SetPrefix("_test_")
	err = handler.LogEntry(NewLogEntry(0, LevelInfo, "hello two"))
	assert.NoError(t, err)

	assert.Equal(t, "hello one\n_test_hello two\n", buffer.String())
}

func ExampleConsoleHandler() {
	logger := NewLogger(NewConsoleHandler("example ", 0))
	logger.Info("hello world")
	// Output: example hello world
}
