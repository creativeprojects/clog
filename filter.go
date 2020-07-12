package clog

// LevelFilter is a log middleware that is only passing log entries of level >= minimum level
type LevelFilter struct {
	destLog  Handler
	minLevel LogLevel
}

// NewLevelFilter creates a new LevelFilter handler
// passing log entries to destination if level >= minimum level
func NewLevelFilter(minLevel LogLevel, destination Handler) *LevelFilter {
	return &LevelFilter{
		minLevel: minLevel,
		destLog:  destination,
	}
}

// SetLevel changes the minimum level the log entries are going to be sent to the destination logger
func (l *LevelFilter) SetLevel(minLevel LogLevel) {
	l.minLevel = minLevel
}

// Log the LogEntry
func (l *LevelFilter) Log(logEntry LogEntry) error {
	if logEntry.Level < l.minLevel {
		return nil
	}
	l.destLog.Log(logEntry)
	return nil
}

// Verify interface
var (
	_ Handler = &LevelFilter{}
)
