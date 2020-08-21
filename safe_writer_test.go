package clog

import (
	"bytes"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeWriterConcurrency(t *testing.T) {
	iterations := 10000
	buffer := &bytes.Buffer{}
	writer := NewSafeWriter(buffer)
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
