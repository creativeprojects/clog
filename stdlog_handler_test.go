package clog

import (
	"bytes"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrefix(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "prefix", 0)
	err := handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"message"},
	})
	assert.NoError(t, err)
	assert.Equal(t, "prefixDEBUG message\n", buffer.String())
}

func TestSetPrefix(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "no prefix", 0)
	handler.SetPrefix("prefix")
	err := handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"message"},
	})
	assert.NoError(t, err)
	assert.Equal(t, "prefixDEBUG message\n", buffer.String())
}

func TestStandardLogHandlerConcurrency(t *testing.T) {
	iterations := 1000
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "", 0)
	logger := NewLogger(handler)
	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func(num int) {
			logger.Infof("log %03d", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for line, err := buffer.ReadString('\n'); err == nil; line, err = buffer.ReadString('\n') {
		assert.Len(t, line, 14)
	}
}

func TestCannotClose(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "", 0)
	err := handler.Close()
	assert.Error(t, err)
}

func TestCanClose(t *testing.T) {
	file, err := os.CreateTemp(t.TempDir(), "TestCanClose")
	require.NoError(t, err)
	handler := NewStandardLogHandler(file, "", 0)
	err = handler.Close()
	assert.NoError(t, err)
}

func TestChangeOutput(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(os.Stdout, "", 0)
	handler.SetOutput(buffer)
	err := handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"message"},
	})
	assert.NoError(t, err)
	assert.Equal(t, "DEBUG message\n", buffer.String())
}

func TestLogOnClosedOutput(t *testing.T) {
	logFile := filepath.Join(t.TempDir(), "file.log")
	file, err := os.Create(logFile)
	require.NoError(t, err)

	handler := NewStandardLogHandler(file, "", 0)
	err = handler.Close()
	assert.NoError(t, err)

	err = handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"message"},
	})
	assert.Error(t, err)
}
