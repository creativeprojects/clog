package clog

// TeeHandler sends logs to two handlers at the same time
type TeeHandler struct {
	firstHandler  Handler
	secondHandler Handler
}

// NewTeeHandler creates a handler that redirects logs to 2 handlers at the same time
func NewTeeHandler(primary, backup Handler) *TeeHandler {
	return &TeeHandler{
		firstHandler:  primary,
		secondHandler: backup,
	}
}

// LogEntry send messages to both handlers
func (h *TeeHandler) LogEntry(logEntry LogEntry) error {
	if h.firstHandler == nil {
		return ErrNoPrimaryHandler
	}
	// don't wait until we get an error to also check the backup handler
	if h.secondHandler == nil {
		return ErrNoBackupHandler
	}
	logEntry.Calldepth++
	err := h.firstHandler.LogEntry(logEntry)
	if err != nil {
		return err
	}
	return h.secondHandler.LogEntry(logEntry)
}

// SetPrefix sets a prefix on every log message
func (h *TeeHandler) SetPrefix(prefix string) Handler {
	if h.firstHandler != nil {
		prefixer, ok := h.firstHandler.(Prefixer)
		if ok {
			prefixer.SetPrefix(prefix)
		}
	}
	if h.secondHandler != nil {
		prefixer, ok := h.secondHandler.(Prefixer)
		if ok {
			prefixer.SetPrefix(prefix)
		}
	}
	return h
}

// Verify interface
var (
	_ Handler = &TeeHandler{}
)
