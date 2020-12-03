package clog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type asyncTestHandler struct {
	received int
}

func (h *asyncTestHandler) LogEntry(LogEntry) error {
	h.received++
	time.Sleep(10 * time.Microsecond)
	return nil
}
func (h *asyncTestHandler) SetPrefix(string) Handler {
	return h
}

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

func TestAsyncHandlerCanCanSetPrefix(t *testing.T) {
	handler := NewMemoryHandler()
	async := NewAsyncHandler(handler)
	async.SetPrefix("_test_")
	err := async.LogEntry(NewLogEntry(3, LevelInfo, "hello world"))
	assert.NoError(t, err)
	// wait for the logs to be written
	async.Close()
	assert.Equal(t, "_test_hello world", handler.log[0])
}

func TestCannotSendLogsToClosedAsyncHandler(t *testing.T) {
	handler := NewMemoryHandler()
	async := NewAsyncHandler(handler)
	err := async.LogEntry(NewLogEntry(3, LevelInfo, "hello world"))
	assert.NoError(t, err)

	async.Close()
	err = async.LogEntry(NewLogEntry(3, LevelInfo, "hello world"))
	assert.Error(t, err)
}

func TestAsyncHandlerSetNextHandler(t *testing.T) {
	memHandler := NewMemoryHandler()
	handler := NewAsyncHandler(memHandler)
	assert.NotNil(t, handler.GetHandler())

	handler.SetHandler(nil)
	assert.Nil(t, handler.GetHandler())
}

func ExampleAsyncHandler() {
	// Close the async handler before finishing your application or you might miss the last few log messages
	handler := NewAsyncHandler(NewTextHandler("async ", 0))
	defer handler.Close()

	logger := NewLogger(handler)
	logger.Info("hello world")
	// Output: async hello world
}
