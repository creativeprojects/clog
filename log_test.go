package clog

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkStreamMessages(b *testing.B) {
	b.ReportAllocs()
	streamHandler := NewStandardLogHandler(ioutil.Discard, "", log.LstdFlags)
	logger := NewLogger(NewLevelFilter(LevelDebug, streamHandler))
	param1 := "string"
	param2 := 0

	for i := 0; i < b.N; i++ {
		logger.Info("Message with", param1, param2)
	}
}

func BenchmarkStreamFilteredMessages(b *testing.B) {
	b.ReportAllocs()
	streamHandler := NewStandardLogHandler(ioutil.Discard, "", log.LstdFlags)
	logger := NewLogger(NewLevelFilter(LevelWarning, streamHandler))
	param1 := "string"
	param2 := 0

	for i := 0; i < b.N; i++ {
		logger.Info("Message with", param1, param2)
	}
}

func BenchmarkStreamFormattedMessages(b *testing.B) {
	b.ReportAllocs()
	streamHandler := NewStandardLogHandler(ioutil.Discard, "", log.LstdFlags)
	logger := NewLogger(NewLevelFilter(LevelDebug, streamHandler))
	param1 := "string"
	param2 := 0

	for i := 0; i < b.N; i++ {
		logger.Infof("Message with a %s and a %d", param1, param2)
	}
}

func BenchmarkStreamFilteredFormattedMessages(b *testing.B) {
	b.ReportAllocs()
	streamHandler := NewStandardLogHandler(ioutil.Discard, "", log.LstdFlags)
	logger := NewLogger(NewLevelFilter(LevelWarning, streamHandler))
	param1 := "string"
	param2 := 0

	for i := 0; i < b.N; i++ {
		logger.Infof("Message with a %s and a %d", param1, param2)
	}
}

func TestDefaultHandler(t *testing.T) {
	// Initial type
	defer SetDefaultLogger(nil)
	assert.IsType(t, new(overflowHandler), GetDefaultLogger().GetHandler())

	// Custom type
	SetDefaultLogger(NewConsoleLogger())
	assert.IsType(t, new(ConsoleHandler), GetDefaultLogger().GetHandler())

	// Overflow behaviour
	SetDefaultLogger(nil)
	require.IsType(t, new(overflowHandler), GetDefaultLogger().GetHandler())

	handler := GetDefaultLogger().GetHandler().(*overflowHandler)
	assert.Equal(t, defaultLogBufferSize, handler.overflowSize)
	assert.True(t, handler.Empty())

	for handler == GetDefaultLogger().GetHandler() {
		Info("log line")
	}
	assert.IsType(t, new(DiscardHandler), GetDefaultLogger().GetHandler())
}

func TestSetDefaultHandler(t *testing.T) {
	// Initial type
	defer SetDefaultLogger(nil)

	SetDefaultLogger(nil)
	Info("log line")

	handler := NewMemoryHandler()
	assert.True(t, handler.Empty())

	GetDefaultLogger().SetHandler(handler)
	assert.Same(t, handler, GetDefaultLogger().GetHandler())
	assert.Equal(t, []string{"log line"}, handler.Logs())
}

func TestSetDefaultLogger(t *testing.T) {
	// Initial type
	defer SetDefaultLogger(nil)

	SetDefaultLogger(nil)
	Info("log line")

	handler := NewMemoryHandler()
	logger := NewLogger(handler)
	assert.True(t, handler.Empty())

	SetDefaultLogger(logger)
	assert.Same(t, logger, GetDefaultLogger())
	assert.Same(t, handler, GetDefaultLogger().GetHandler())
	assert.Equal(t, []string{"log line"}, handler.Logs())
}

func TestPackage(t *testing.T) {
	SetDefaultLogger(NewConsoleLogger())
	defer SetDefaultLogger(nil)

	buffer := &strings.Builder{}
	handler := defaultLogger.GetHandler().(*ConsoleHandler)
	handler.Colouring(false)
	handler.logger.SetOutput(buffer)
	Trace("trace")
	Tracef("%s", "trace")
	Debug("debug")
	Debugf("%s", "debug")
	Info("info")
	Infof("%s", "info")
	Warning("warning")
	Warningf("%s", "warning")
	Error("error")
	Errorf("%s", "error")
	assert.Equal(t, "trace\ntrace\ndebug\ndebug\ninfo\ninfo\nwarning\nwarning\nerror\nerror\n", buffer.String())
}
