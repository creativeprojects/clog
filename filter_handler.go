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
func (l *LevelFilter) SetLevel(minLevel LogLevel) {
	l.minLevel = minLevel
}

// SetHandler sets a new handler for the filter
func (l *LevelFilter) SetHandler(handler Handler) {
	l.handler = handler
}

// GetHandler returns the current handler used by the filter
func (l *LevelFilter) GetHandler() Handler {
	return l.handler
}

// LogEntry the LogEntry
func (l *LevelFilter) LogEntry(logEntry LogEntry) error {
	if logEntry.Level < l.minLevel {
		return nil
	}
	logEntry.Calldepth++
	return l.handler.LogEntry(logEntry)
}

// Verify interface
var (
	_ Handler = &LevelFilter{}
)
