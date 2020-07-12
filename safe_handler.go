package clog

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

// Log send messages to primaryHandler first, then to backupHandler if it had fail
func (l *SafeHandler) Log(logEntry LogEntry) error {
	err := l.primaryHandler.Log(logEntry)
	if err != nil {
		return l.backupHandler.Log(logEntry)
	}
	return nil
}

// Verify interface
var (
	_ Handler = &SafeHandler{}
)
