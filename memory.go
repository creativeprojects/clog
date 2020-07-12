package clog

// MemoryHandler save messages in memory (useful for unit test).
type MemoryHandler struct {
	log []string
}

// NewMemoryHandler creates a new MemoryHandler that keep log in memory.
func NewMemoryHandler() *MemoryHandler {
	return &MemoryHandler{
		log: make([]string, 0, 10),
	}
}

// Log keep the messages in memory.
func (l *MemoryHandler) Log(logEntry LogEntry) error {
	l.log = append(l.log, logEntry.GetMessage())
	return nil
}

// Verify interface
var (
	_ Handler = &MemoryHandler{}
)
