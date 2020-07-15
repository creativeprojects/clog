package clog

import (
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
