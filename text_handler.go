package clog

import (
	"log"
	"os"
)

// TextHandler logs messages directly to the console
type TextHandler struct {
	logger *log.Logger
}

// NewTextHandler creates a new handler to send logs to the console
func NewTextHandler(prefix string, flag int) *TextHandler {
	handler := &TextHandler{
		logger: log.New(os.Stdout, prefix, flag),
	}
	return handler
}

// LogEntry sends a log entry with the specified level
func (l *TextHandler) LogEntry(logEntry LogEntry) error {
	return l.logger.Output(logEntry.Calldepth+2, l.levelPrefix(logEntry.Level)+logEntry.GetMessage())
}

// SetPrefix sets a prefix on every log message
func (l *TextHandler) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}

func (l *TextHandler) levelPrefix(logLevel LogLevel) string {
	switch logLevel {
	case LevelError:
		return "Error: "
	case LevelWarning:
		return "Warning: "
	default:
		return ""
	}
}

// Verify interface
var (
	_ Handler = &TextHandler{}
)
