package clog

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	handler := NewMemoryHandler()
	stdlog := log.New(NewWriter(LevelInfo, handler), "prefix: ", 0)
	stdlog.Printf("%s %d", "one", 2)
	assert.Equal(t, []string{"prefix: one 2\n"}, handler.log)
}

func BenchmarkWriter(b *testing.B) {
	b.ReportAllocs()
	stdlog := log.New(NewWriter(LevelInfo, NewDiscardHandler()), "1234567890", 0)

	for i := 0; i < b.N; i++ {
		stdlog.Printf("%s%s", "12345", "12345")
	}
}
