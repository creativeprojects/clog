package clog

var (
	// errorDiscarded is sent when using the Discard handler
	errorDiscarded = ErrMessageDiscarded
)

// DiscardHandler forgets any log message
type DiscardHandler struct{}

// NewDiscardHandler returns a handler that forgets all the logs you throw at it.
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// LogEntry discards the LogEntry
func (l *DiscardHandler) LogEntry(LogEntry) error {
	return errorDiscarded
}

// SetPrefix sets a prefix on every log message
func (l *DiscardHandler) SetPrefix(string) {}

// Verify interface
var (
	_ Handler = &DiscardHandler{}
)
