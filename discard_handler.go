package clog

import "errors"

// DiscardHandler forgets any log message
type DiscardHandler struct{}

// LogEntry discards the LogEntry
func (l *DiscardHandler) LogEntry(LogEntry) error {
	return errors.New("this message is not going anywhere")
}

// Verify interface
var (
	_ Handler = &DiscardHandler{}
)
