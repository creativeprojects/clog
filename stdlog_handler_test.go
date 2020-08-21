package clog

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "prefix", log.Lmsgprefix)
	handler.LogEntry(LogEntry{
		Values: []interface{}{"message"},
	})
	assert.Equal(t, "prefixDEBUG message\n", buffer.String())
}

func TestStandardLogPrefixBefore(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "prefix", 0)
	logger.Print("message")
	assert.Equal(t, "prefixmessage\n", buffer.String())
}

func TestStandardLogPrefixAfter(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "prefix", log.Lmsgprefix)
	logger.Print("message")
	assert.Equal(t, "prefixmessage\n", buffer.String())
}
