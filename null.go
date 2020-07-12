package clog

import "errors"

// NullHandler forgets any log message
type NullHandler struct{}

// Log discards the LogEntry
func (l *NullHandler) Log(LogEntry) error {
	return errors.New("this message is not going anywhere")
}

// Verify interface
var (
	_ Handler = &NullHandler{}
)
