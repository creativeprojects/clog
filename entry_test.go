package clog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogEntry(t *testing.T) {
	entry := NewLogEntry(2, LevelInfo, "a", "b", 10, true, false)
	assert.Equal(t, "ab10 true false", entry.GetMessage())
}

func TestNewLogEntryf(t *testing.T) {
	entry := NewLogEntryf(2, LevelInfo, "%s%s%d %v %v", "a", "b", 10, true, false)
	assert.Equal(t, "ab10 true false", entry.GetMessage())
}
