package clog

// LogEntry represents a log entry
type LogEntry struct {
	Level  LogLevel
	Format string
	Values []interface{}
}
