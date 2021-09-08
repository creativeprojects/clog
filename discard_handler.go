package clog

// DiscardHandler forgets any log message
type DiscardHandler struct{}

// NewDiscardHandler returns a handler that forgets all the logs you throw at it.
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// LogEntry discards the LogEntry
func (h *DiscardHandler) LogEntry(LogEntry) error {
	return ErrMessageDiscarded
}

// Verify interface
var (
	_ Handler = &DiscardHandler{}
)
