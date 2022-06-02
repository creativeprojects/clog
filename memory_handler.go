package clog

import "sync"

// MemoryHandler save messages in memory (useful for unit testing).
type MemoryHandler struct {
	prefix string
	log    []string
	mu     sync.Mutex
}

// NewMemoryHandler creates a new MemoryHandler that keeps logs in memory.
func NewMemoryHandler() *MemoryHandler {
	return &MemoryHandler{
		log: make([]string, 0, 10),
	}
}

// LogEntry keep the message in memory.
func (h *MemoryHandler) LogEntry(logEntry LogEntry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.log = append(h.log, h.prefix+logEntry.GetMessage())
	return nil
}

// SetPrefix adds a prefix to every log message.
// Please note no space is added between the prefix and the log message
func (h *MemoryHandler) SetPrefix(prefix string) Handler {
	h.prefix = prefix
	return h
}

// Logs return a list of all the messages sent to the logger
func (h *MemoryHandler) Logs() []string {
	return h.log
}

// Verify interface
var (
	_ Handler = &MemoryHandler{}
)
