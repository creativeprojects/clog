package clog

import (
	"fmt"
)

// LogEntry represents a log entry
type LogEntry struct {
	Calldepth int
	Level     LogLevel
	Format    string
	Values    []interface{}
}

// GetMessage returns the formatted message from Format & Values
func (l LogEntry) GetMessage() string {
	if l.Format == "" {
		return fmt.Sprint(l.Values...)
	}
	return fmt.Sprintf(l.Format, l.Values...)
}

// GetMessageWithLevelPrefix returns the formatted message from Format & Values prefixed with the level name
func (l LogEntry) GetMessageWithLevelPrefix() string {
	if l.Format == "" {
		return l.Level.String() + " " + fmt.Sprint(l.Values...)
	}
	return fmt.Sprintf(l.Level.String()+" "+l.Format, l.Values...)
}
