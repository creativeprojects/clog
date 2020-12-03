package clog

// LevelFilter is a log middleware that is only passing log entries of level >= minimum level
type LevelFilter struct {
	handler  Handler
	minLevel LogLevel
}

// NewLevelFilter creates a new LevelFilter handler
// passing log entries to destination if level >= minimum level
func NewLevelFilter(minLevel LogLevel, destination Handler) *LevelFilter {
	return &LevelFilter{
		minLevel: minLevel,
		handler:  destination,
	}
}

// SetLevel changes the minimum level the log entries are going to be sent to the destination logger
func (h *LevelFilter) SetLevel(minLevel LogLevel) *LevelFilter {
	h.minLevel = minLevel
	return h
}

// SetHandler sets a new handler for the filter
func (h *LevelFilter) SetHandler(handler Handler) {
	h.handler = handler
}

// GetHandler returns the current handler used by the filter
func (h *LevelFilter) GetHandler() Handler {
	return h.handler
}

// SetPrefix sets a prefix on every log message
func (h *LevelFilter) SetPrefix(prefix string) Handler {
	if h.handler == nil {
		return h
	}
	h.handler.SetPrefix(prefix)
	return h
}

// LogEntry the LogEntry
func (h *LevelFilter) LogEntry(logEntry LogEntry) error {
	if h.handler == nil {
		return ErrNoRegisteredHandler
	}
	if logEntry.Level < h.minLevel {
		return nil
	}
	logEntry.Calldepth++
	return h.handler.LogEntry(logEntry)
}

// Verify interface
var (
	_ Handler           = &LevelFilter{}
	_ MiddlewareHandler = &LevelFilter{}
)
