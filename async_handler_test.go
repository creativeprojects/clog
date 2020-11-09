package clog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsyncHandler(t *testing.T) {
	items := 100000
	handler := NewMemoryHandler()
	asyncHandler := NewAsyncHandler(handler)
	logger := NewLogger(asyncHandler)
	for i := 0; i < items; i++ {
		logger.Infof("message %d", i)
	}
	start := time.Now()
	asyncHandler.Wait()
	t.Logf("was waiting for %s for all the messages to be sent", time.Since(start))
	assert.Len(t, handler.log, items)
}

func TestAsyncHandlerShouldFail(t *testing.T) {
	handler := NewAsyncHandler(nil)
	err := handler.LogEntry(LogEntry{})
	assert.Error(t, err)
}
