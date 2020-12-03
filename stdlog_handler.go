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

// LogEntry sends a log entry with the specified level.
func (h *StandardLogHandler) LogEntry(logEntry LogEntry) error {
	return h.stdlog.Output(logEntry.Calldepth+2, logEntry.GetMessageWithLevelPrefix())
}

// SetOutput sets the output destination for the logger.
func (h *StandardLogHandler) SetOutput(output io.Writer) *StandardLogHandler {
	h.stdlog.SetOutput(output)
	return h
}

// SetPrefix sets a prefix on every log message
func (h *StandardLogHandler) SetPrefix(prefix string) Handler {
	h.stdlog.SetPrefix(prefix)
	return h
}

// Verify interface
var (
	_ Handler = &StandardLogHandler{}
)
