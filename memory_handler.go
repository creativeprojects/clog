package clog

import "sync"

// MemoryHandler save messages in memory (useful for unit testing).
type MemoryHandler struct {
	log []string
	mu  sync.Mutex
}

// NewMemoryHandler creates a new MemoryHandler that keeps logs in memory.
func NewMemoryHandler() *MemoryHandler {
	return &MemoryHandler{
		log: make([]string, 0, 10),
	}
}

// LogEntry keep the message in memory.
func (l *MemoryHandler) LogEntry(logEntry LogEntry) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.log = append(l.log, logEntry.GetMessage())
	return nil
}

// SetPrefix does nothing on the memory handler
func (l *MemoryHandler) SetPrefix(prefix string) {}

// Verify interface
var (
	_ Handler = &MemoryHandler{}
)
