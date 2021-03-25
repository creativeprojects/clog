package clog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterChainLogger(t *testing.T) {
	expected1 := []string{
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
	expected2 := []string{
		"TRACE 0 < 1",
		"TRACE 0 < 2",
		"DEBUG 1 < 2",
		"TRACE 0 < 3",
		"DEBUG 1 < 3",
		"INFO  2 < 3",
		"TRACE 0 < 4",
		"DEBUG 1 < 4",
		"INFO  2 < 4",
		"WARN  3 < 4",
	}

	buffer1, buffer2 := &bytes.Buffer{}, &bytes.Buffer{}
	handler := NewLevelFilterChain(LevelTrace, NewStandardLogHandler(buffer1, "", 0), NewStandardLogHandler(buffer2, "", 0))
	logger := NewLogger(handler)

	for minLevel := LevelTrace; minLevel <= LevelError; minLevel++ {
		handler.SetLevel(minLevel)
		for logLevel := LevelTrace; logLevel <= LevelError; logLevel++ {
			if logLevel >= minLevel {
				logger.Logf(logLevel, "%d >= %d", logLevel, minLevel)
			} else {
				logger.Logf(logLevel, "%d < %d", logLevel, minLevel)
			}
		}
	}
	logs := []string{}
	for line, err := buffer1.ReadString('\n'); err == nil; line, err = buffer1.ReadString('\n') {
		logs = append(logs, strings.Trim(line, "\n"))
	}
	assert.ElementsMatch(t, expected1, logs)

	logs = []string{}
	for line, err := buffer2.ReadString('\n'); err == nil; line, err = buffer2.ReadString('\n') {
		logs = append(logs, strings.Trim(line, "\n"))
	}
	assert.ElementsMatch(t, expected2, logs)
}

func TestFilterChainHandlerShouldFail(t *testing.T) {
	handler := NewLevelFilterChain(LevelInfo, nil, NewTextHandler("", 0))
	err := handler.LogEntry(LogEntry{Level: LevelWarning})
	assert.Error(t, err)
}

func TestFilterChainHandlerShouldFailOnNext(t *testing.T) {
	handler := NewLevelFilterChain(LevelInfo, NewTextHandler("", 0), nil)
	err := handler.LogEntry(LogEntry{Level: LevelDebug})
	assert.Error(t, err)
}

func TestFilterChainHandlerCanChangeHandler(t *testing.T) {
	handler := NewLevelFilterChain(LevelInfo, nil, nil)
	assert.Nil(t, handler.GetHandler())

	next := NewDiscardHandler()
	handler.SetHandler(next)
	assert.Equal(t, next, handler.GetHandler())
}

func TestFilterChainHandlerCanChangeNextHandler(t *testing.T) {
	handler := NewLevelFilterChain(LevelInfo, nil, nil)
	assert.Nil(t, handler.GetNextHandler())

	next := NewDiscardHandler()
	handler.SetNextHandler(next)
	assert.Equal(t, next, handler.GetNextHandler())
}

func TestFilterChainHandlerCanCanSetPrefix(t *testing.T) {
	handler1 := NewMemoryHandler()
	handler2 := NewMemoryHandler()
	filter := NewLevelFilterChain(LevelInfo, handler1, handler2)
	filter.SetPrefix("_test_")

	// send to the first handler
	err := filter.LogEntry(NewLogEntry(3, LevelInfo, "hello world info"))
	assert.NoError(t, err)
	assert.Len(t, handler1.log, 1)
	assert.Len(t, handler2.log, 0)
	assert.Equal(t, "_test_hello world info", handler1.log[0])

	// now send to the next handler
	err = filter.LogEntry(NewLogEntry(3, LevelDebug, "hello world debug"))
	assert.NoError(t, err)
	assert.Len(t, handler1.log, 1)
	assert.Len(t, handler2.log, 1)
	assert.Equal(t, "hello world debug", handler2.log[0])
}

func ExampleLevelFilterChain() {
	logger := NewLogger(
		NewLevelFilterChain(LevelInfo,
			NewTextHandler("info and more: ", 0),
			NewLevelFilterChain(LevelDebug,
				NewTextHandler("debug: ", 0),
				nil)))
	logger.Debug("some debug")
	logger.Info("some info")
	// Output:
	// debug: some debug
	// info and more: some info
}
