package clog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "prefix", 0)
	handler.LogEntry(LogEntry{
		Values: []interface{}{"message"},
	})
	assert.Equal(t, "prefixDEBUG message\n", buffer.String())
}
