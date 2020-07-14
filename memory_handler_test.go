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
	assert.Len(t, handler.log, iterations)
	for _, entry := range handler.log {
		assert.Len(t, entry, 7)
	}
}
