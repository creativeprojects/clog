package clog

import "errors"

// SafeHandler sends logs to an alternate destination when the primary destination fails
type SafeHandler struct {
	primaryHandler Handler
	backupHandler  Handler
}

// NewSafeHandler creates a handler that redirects logs to a backup handler when the primary fails
func NewSafeHandler(primary, backup Handler) *SafeHandler {
	return &SafeHandler{
		primaryHandler: primary,
		backupHandler:  backup,
	}
}

// LogEntry send messages to primaryHandler first, then to backupHandler if it had fail
func (h *SafeHandler) LogEntry(logEntry LogEntry) error {
	if h.primaryHandler == nil {
		return errors.New("no primary handler registered")
	}
	// don't wait until we get an error to also check the backup handler
	if h.backupHandler == nil {
		return errors.New("no backup handler registered")
	}
	logEntry.Calldepth++
	err := h.primaryHandler.LogEntry(logEntry)
	if err != nil {
		return h.backupHandler.LogEntry(logEntry)
	}
	return nil
}

// SetPrefix sets a prefix on every log message
func (h *SafeHandler) SetPrefix(prefix string) {
	if h.primaryHandler != nil {
		h.primaryHandler.SetPrefix(prefix)
	}
	if h.backupHandler != nil {
		h.backupHandler.SetPrefix(prefix)
	}
}

// Verify interface
var (
	_ Handler = &SafeHandler{}
)
