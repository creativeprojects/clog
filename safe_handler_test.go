package clog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleErrorShouldReturnError(t *testing.T) {
	handler := NewSafeHandler(&DiscardHandler{}, &DiscardHandler{})
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}
