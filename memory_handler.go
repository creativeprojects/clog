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
	message := logEntry.GetMessage()
	{
		h.mu.Lock()
		defer h.mu.Unlock()

		h.log = append(h.log, h.prefix+message)
		return nil
	}
}

// SetPrefix adds a prefix to every log message.
// Please note no space is added between the prefix and the log message
func (h *MemoryHandler) SetPrefix(prefix string) Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.prefix = prefix
	return h
}

// Logs return a list of all the messages sent to the logger
func (h *MemoryHandler) Logs() []string {
	h.mu.Lock()
	defer h.mu.Unlock()

	return append([]string(nil), h.log...)
}

// Empty return true when the internal list of logs is empty
func (h *MemoryHandler) Empty() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return len(h.log) == 0
}

// Pop returns the latest log from the internal storage (and removes it)
func (h *MemoryHandler) Pop() string {
	h.mu.Lock()
	defer h.mu.Unlock()

	latest := h.log[len(h.log)-1]
	h.log = h.log[:len(h.log)-1]
	return latest
}

// Verify interface
var (
	_ Handler = &MemoryHandler{}
)
