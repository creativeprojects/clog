package clog

import (
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
