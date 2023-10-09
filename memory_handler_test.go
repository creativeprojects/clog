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
	assert.Equal(t, (memoryLogFixedSize+7)*int64(iterations), handler.Size())
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

	assert.False(t, handler.Empty())
	handler.Pop()
	handler.Pop()
	assert.True(t, handler.Empty())
	assert.Panics(t, func() { handler.Pop() })
}

func TestMemoryHandlerPopWithLevel(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewLogger(handler)
	logger.SetPrefix("prefix:")
	logger.Error("err-msg")
	logger.Info("i2")
	logger.Info("i1")

	assert.Len(t, handler.Logs(), 3)
	assert.Equal(t, "prefix:i1", handler.Pop())

	level, prefix, message := handler.PopWithLevel()
	assert.Equal(t, LevelInfo, level)
	assert.Equal(t, "prefix:", prefix)
	assert.Equal(t, "i2", message)

	level, prefix, message = handler.PopWithLevel()
	assert.Equal(t, LevelError, level)
	assert.Equal(t, "prefix:", prefix)
	assert.Equal(t, "err-msg", message)
}

func TestMemoryHandlerReset(t *testing.T) {
	handler := NewMemoryHandler()
	NewLogger(handler).Info("info")
	assert.False(t, handler.Empty())
	handler.Reset()
	assert.True(t, handler.Empty())
}

func TestMemoryHandlerTransferTo(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewLogger(handler)
	for i := 0; i < 3; i++ {
		logger.Warningf("log %d", i)
	}

	assert.Len(t, handler.Logs(), 3)
	assert.False(t, handler.TransferTo(handler))

	dest := NewMemoryHandler()
	for i := 0; i < 2; i++ {
		assert.Equal(t, i == 0, handler.TransferTo(dest))
		assert.Len(t, handler.Logs(), 0)
		assert.Len(t, dest.Logs(), 3)
	}

	assert.Panics(t, func() { handler.Pop() })
	assert.Equal(t, []string{"log 0", "log 1", "log 2"}, dest.Logs())

	level, _, message := dest.PopWithLevel()
	assert.Equal(t, LevelWarning, level)
	assert.Equal(t, "log 2", message)

	assert.Len(t, dest.Logs(), 2)
	assert.True(t, dest.TransferTo(nil))
	assert.Empty(t, dest.Logs())
}

func TestMemoryHandlerTransferToPrefix(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewLogger(handler)
	logger.SetPrefix("prefix ")
	for i := 0; i < 3; i++ {
		logger.Warningf("log %d", i)
	}

	assert.Equal(t, "prefix log 2", handler.Pop())

	// TransferTo strips logger specific prefix
	receiver := NewMemoryHandler()
	handler.TransferTo(receiver)
	assert.Equal(t, []string{"log 0", "log 1"}, receiver.Logs())
}
