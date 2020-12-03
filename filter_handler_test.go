package clog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilteredLogger(t *testing.T) {
	expected := []string{
		"TRACE 0 >= 0",
		"DEBUG 1 >= 0",
		"INFO  2 >= 0",
		"WARN  3 >= 0",
		"ERROR 4 >= 0",
		"DEBUG 1 >= 1",
		"INFO  2 >= 1",
		"WARN  3 >= 1",
		"ERROR 4 >= 1",
		"INFO  2 >= 2",
		"WARN  3 >= 2",
		"ERROR 4 >= 2",
		"WARN  3 >= 3",
		"ERROR 4 >= 3",
		"ERROR 4 >= 4",
	}

	buffer := &bytes.Buffer{}
	handler := NewLevelFilter(LevelTrace, NewStandardLogHandler(buffer, "", 0))
	logger := NewLogger(handler)

	for minLevel := LevelTrace; minLevel <= LevelError; minLevel++ {
		handler.SetLevel(minLevel)
		for logLevel := LevelTrace; logLevel <= LevelError; logLevel++ {
			logger.Logf(logLevel, "%d >= %d", logLevel, minLevel)
		}
	}
	logs := []string{}
	for line, err := buffer.ReadString('\n'); err == nil; line, err = buffer.ReadString('\n') {
		logs = append(logs, strings.Trim(line, "\n"))
	}
	assert.ElementsMatch(t, expected, logs)
}

func TestFilterHandlerShouldFail(t *testing.T) {
	handler := NewLevelFilter(LevelInfo, nil)
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}

func TestFilterHandlerCanChangeHandler(t *testing.T) {
	handler := NewLevelFilter(LevelInfo, nil)
	assert.Nil(t, handler.GetHandler())

	next := NewDiscardHandler()
	handler.SetHandler(next)
	assert.Equal(t, next, handler.GetHandler())
}

func TestFilterHandlerCanCanSetPrefix(t *testing.T) {
	handler := NewMemoryHandler()
	filter := NewLevelFilter(LevelInfo, handler)
	filter.SetPrefix("_test_")
	err := filter.LogEntry(NewLogEntry(3, LevelInfo, "hello world"))
	assert.NoError(t, err)
	assert.Equal(t, "_test_hello world", handler.log[0])
}

func ExampleLevelFilter() {
	logger := NewLogger(NewLevelFilter(LevelInfo, NewTextHandler("", 0)))
	logger.Debug("won't be displayed")
	logger.Info("hello world")
	// Output: hello world
}
