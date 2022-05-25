package clog

// TeeHandler sends logs to two handlers at the same time
type TeeHandler struct {
	firstHandler  Handler
	secondHandler Handler
}

// NewTeeHandler creates a handler that redirects logs to 2 handlers at the same time
func NewTeeHandler(a, b Handler) *TeeHandler {
	return &TeeHandler{
		firstHandler:  a,
		secondHandler: b,
	}
}

// LogEntry send messages to both handlers
func (h *TeeHandler) LogEntry(logEntry LogEntry) error {
	var err1, err2 error
	logEntry.Calldepth++
	if h.firstHandler != nil {
		err1 = h.firstHandler.LogEntry(logEntry)
	} else {
		err1 = ErrNoRegisteredHandler
	}
	if h.secondHandler != nil {
		err2 = h.secondHandler.LogEntry(logEntry)
	} else {
		err2 = ErrNoRegisteredHandler
	}
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
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
