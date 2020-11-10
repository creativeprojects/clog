package clog

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextHandler(t *testing.T) {
	handler := NewTextHandler("", 0)
	for level := LevelDebug; level <= LevelError; level++ {
		err := handler.LogEntry(LogEntry{
			Level:  level,
			Values: []interface{}{level},
		})
		assert.NoError(t, err)
	}
}

func TestTextHandlerPrefix(t *testing.T) {
	buffer := &strings.Builder{}
	handler := NewTextHandler("", 0)
	// manually change the output to our local buffer
	handler.logger.SetOutput(buffer)

	handler.LogEntry(NewLogEntry(0, LevelInfo, "hello one"))
	handler.SetPrefix("_test_")
	handler.LogEntry(NewLogEntry(0, LevelInfo, "hello two"))

	assert.Equal(t, "hello one\n_test_hello two\n", buffer.String())
}
