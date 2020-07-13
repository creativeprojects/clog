package clog

import (
	"testing"
)

func BenchmarkDiscardHandler(b *testing.B) {
	b.ReportAllocs()
	logger := NewLogger(&DiscardHandler{})
	param1 := "string"
	param2 := 0

	for i := 0; i < b.N; i++ {
		logger.Info("Message with", param1, param2)
	}
}

func BenchmarkDiscardHandlerStaticMessage(b *testing.B) {
	b.ReportAllocs()
	logger := NewLogger(&DiscardHandler{})

	for i := 0; i < b.N; i++ {
		logger.Info("message")
	}
}
