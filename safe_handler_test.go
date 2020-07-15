package clog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorShouldSendToBackup(t *testing.T) {
	backup := NewMemoryHandler()
	handler := NewSafeHandler(&DiscardHandler{}, backup)
	err := handler.LogEntry(LogEntry{})
	assert.NoError(t, err)
	assert.Len(t, backup.log, 1)
}

func TestDoubleErrorShouldReturnError(t *testing.T) {
	handler := NewSafeHandler(&DiscardHandler{}, &DiscardHandler{})
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}
