package clog

import (
	"bytes"
	"io/ioutil"
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileLoggerConcurrency(t *testing.T) {

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
