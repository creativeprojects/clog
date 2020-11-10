package clog

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	buffer := &strings.Builder{}
	handler := NewStandardLogHandler(buffer, "", 0)
	logger := NewLogger(handler)
	logger.Trace("trace")
	logger.Tracef("%s", "trace")
	logger.Debug("debug")
	logger.Debugf("%s", "debug")
	logger.Info("info")
	logger.Infof("%s", "info")
	logger.Warning("warning")
	logger.Warningf("%s", "warning")
	logger.Error("error")
	logger.Errorf("%s", "error")

	assert.Equal(t, `TRACE trace
TRACE trace
DEBUG debug
DEBUG debug
INFO  info
INFO  info
WARN  warning
WARN  warning
ERROR error
ERROR error
`, buffer.String())
}

func TestLoggerShouldFail(t *testing.T) {
	handler := NewLogger(nil)
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}

func TestLoggerSetPrefix(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewLogger(handler)
	logger.SetPrefix("_test_")
	logger.Debug("hello world")
	assert.Equal(t, "_test_hello world", handler.log[0])
}

func TestLoggerSetHandler(t *testing.T) {
	logger := NewFilteredConsoleLogger(LevelInfo)
	assert.NotNil(t, logger.GetHandler())

	logger.SetHandler(nil)
	assert.Nil(t, logger.GetHandler())
}
