package clog

import (
	"io/ioutil"
	"log"
	"testing"
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
