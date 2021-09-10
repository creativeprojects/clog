package clog

import (
	"bytes"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncWriterConcurrency(t *testing.T) {
	iterations := 5000
	buffer := &bytes.Buffer{}
	writer := NewSyncWriter(buffer)
	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func(num int) {
			n, err := writer.Write([]byte("123456789\n"))
			assert.Equal(t, 10, n)
			assert.NoError(t, err)
			wg.Done()
		}(i)
	}
	wg.Wait()
	lines := 0
	for line, err := buffer.ReadString('\n'); err == nil; line, err = buffer.ReadString('\n') {
		assert.Len(t, line, 10)
		lines++
	}
	assert.Equal(t, iterations, lines)
}

type fdBuffer struct {
	bytes.Buffer
}

func (b fdBuffer) Fd() uintptr {
	return 0
}

func TestSyncFdWriterConcurrency(t *testing.T) {
	iterations := 5000
	buffer := &fdBuffer{}
	writer := NewSyncWriter(buffer)
	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func(num int) {
			n, err := writer.Write([]byte("123456789\n"))
			assert.Equal(t, 10, n)
			assert.NoError(t, err)
			wg.Done()
		}(i)
	}
	wg.Wait()
	lines := 0
	for line, err := buffer.ReadString('\n'); err == nil; line, err = buffer.ReadString('\n') {
		assert.Len(t, line, 10)
		lines++
	}
	assert.Equal(t, iterations, lines)
}

func TestSyncWriterNoFd(t *testing.T) {
	_, ok := NewSyncWriter(&bytes.Buffer{}).(interface {
		Fd() uintptr
	})

	if ok {
		t.Error("NewSyncWriter should not expose a Fd method")
	}
}

func TestSyncWriterFd(t *testing.T) {
	_, ok := NewSyncWriter(os.Stdout).(interface {
		Fd() uintptr
	})

	if !ok {
		t.Error("NewSyncWriter does not pass through Fd method")
	}
}
