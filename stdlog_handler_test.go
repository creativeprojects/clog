package clog

import (
	"bytes"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "prefix", 0)
	handler.LogEntry(LogEntry{
		Values: []interface{}{"message"},
	})
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
