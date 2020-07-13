package clog

import (
	"fmt"
)

// LogEntry represents a log entry
type LogEntry struct {
	Calldepth int           // Calldepth is used to calculate the right place where we called the log method
	Level     LogLevel      // Debug, Info, Warning or Error
	Format    string        // Format for *printf (leave blank for *print)
	Values    []interface{} // Values for *print and *printf
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
