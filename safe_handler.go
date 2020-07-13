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

// LogEntry send messages to primaryHandler first, then to backupHandler if it had fail
func (l *SafeHandler) LogEntry(logEntry LogEntry) error {
	logEntry.Calldepth++
	err := l.primaryHandler.LogEntry(logEntry)
	if err != nil {
		return l.backupHandler.LogEntry(logEntry)
	}
	return nil
}

// Verify interface
var (
	_ Handler = &SafeHandler{}
)
