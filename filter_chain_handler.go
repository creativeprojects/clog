package clog

// LevelFilterChain is a log middleware that is only passing log entries of level >= minimum level
type LevelFilterChain struct {
	handler  Handler
	minLevel LogLevel
	next     Handler
}

// NewLevelFilterChain creates a new LevelFilterChain handler
// passing log entries to destination handler if level >= minimum level,
// and passing them to next handler if level < minimum level
func NewLevelFilterChain(minLevel LogLevel, destination, next Handler) *LevelFilterChain {
	return &LevelFilterChain{
		minLevel: minLevel,
		handler:  destination,
		next:     next,
	}
}

// SetLevel changes the minimum level the log entries are going to be sent to the destination logger
func (h *LevelFilterChain) SetLevel(minLevel LogLevel) *LevelFilterChain {
	h.minLevel = minLevel
	return h
}

// SetHandler sets a new handler for the filter
func (h *LevelFilterChain) SetHandler(handler Handler) {
	h.handler = handler
}

// GetHandler returns the current handler used by the filter
func (h *LevelFilterChain) GetHandler() Handler {
	return h.handler
}

// SetNextHandler sets the next handler in chain
func (h *LevelFilterChain) SetNextHandler(handler Handler) {
	h.next = handler
}

// GetNextHandler returns the next handler in chain
func (h *LevelFilterChain) GetNextHandler() Handler {
	return h.next
}

// SetPrefix sets a prefix on every log message
func (h *LevelFilterChain) SetPrefix(prefix string) Handler {
	if h.handler == nil {
		return h
	}
	prefixer, ok := h.handler.(Prefixer)
	if ok {
		prefixer.SetPrefix(prefix)
	}
	return h
}

// LogEntry the LogEntry
func (h *LevelFilterChain) LogEntry(logEntry LogEntry) error {
	if logEntry.Level < h.minLevel {
		if h.next == nil {
			return ErrNoRegisteredHandler
		}
		return h.next.LogEntry(logEntry)
	}
	if h.handler == nil {
		return ErrNoRegisteredHandler
	}
	logEntry.Calldepth++
	return h.handler.LogEntry(logEntry)
}

// Verify interface
var (
	_ Handler           = &LevelFilterChain{}
	_ MiddlewareHandler = &LevelFilterChain{}
)
