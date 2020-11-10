package clog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type asyncTestHandler struct {
	received int
}

func (l *asyncTestHandler) LogEntry(LogEntry) error {
	l.received++
	time.Sleep(10 * time.Microsecond)
	return nil
}
func (l *asyncTestHandler) SetPrefix(string) {}

// test three stages:
// - fill in the buffered channel (3 times)
// - sleep enough time to drain the channel (twice)
// - on the last round, call the Close method which is waiting for the buffered channel to be drained
func TestAsyncHandler(t *testing.T) {
	items := 10000
	loops := 3
	handler := &asyncTestHandler{}
	asyncHandler := NewAsyncHandler(handler)
	logger := NewLogger(asyncHandler)
	for l := 0; l < loops; l++ {
		for i := 0; i < items; i++ {
			logger.Infof("message %d%d", l, i)
		}
		// don't run this on the last loop
		if l == loops-1 {
			break
		}
		before := len(asyncHandler.entries)
		time.Sleep(20 * time.Millisecond)
		t.Logf("channel drained from %d %d to entries", before, len(asyncHandler.entries))
	}
	start := time.Now()
	asyncHandler.Close()
	t.Logf("was waiting for %s for all the messages to be sent", time.Since(start))
	assert.Equal(t, items*loops, handler.received, "Expected %d logs, but found %d", items*loops, handler.received)
}

func TestAsyncHandlerShouldFail(t *testing.T) {
	handler := NewAsyncHandler(nil)
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}
