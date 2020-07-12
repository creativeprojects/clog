package clog

import (
	"io"
	"log"
)

// StandardLogHandler send messages to a io.Writer using the standard logger.
type StandardLogHandler struct {
	stdlog *log.Logger
}

// NewStandardLogHandler creates a handler to send the logs to io.Writer through a standard logger.
func NewStandardLogHandler(out io.Writer, prefix string, flag int) *StandardLogHandler {
	handler := &StandardLogHandler{
		stdlog: log.New(out, prefix, flag),
	}
	return handler
}

// Log sends a log entry with the specified level.
func (l *StandardLogHandler) Log(logEntry LogEntry) error {
	return l.stdlog.Output(2, logEntry.GetMessageWithLevelPrefix())
}

// SetOutput sets the output destination for the logger.
func (l *StandardLogHandler) SetOutput(output io.Writer) {
	l.stdlog.SetOutput(output)
}

// Verify interface
var (
	_ Handler = &StandardLogHandler{}
)
