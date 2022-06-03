package clog

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryHandlerConcurrency(t *testing.T) {
	iterations := 1000
	handler := NewMemoryHandler()
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
	assert.Len(t, handler.Logs(), iterations)
	for _, entry := range handler.Logs() {
		assert.Len(t, entry, 7)
	}
}

func TestMemoryHandlerPop(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewLogger(handler)
	for i := 0; i < 3; i++ {
		logger.Infof("log %d", i)
	}
	assert.Len(t, handler.Logs(), 3)
	log := handler.Pop()
	assert.Len(t, handler.Logs(), 2)
	assert.Equal(t, "log 2", log)
}
