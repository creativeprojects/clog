package clog

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPackage(t *testing.T) {
	buffer := &strings.Builder{}
	handler := defaultLogger.GetHandler().(*ConsoleHandler)
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
