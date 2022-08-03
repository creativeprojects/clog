package clog

import (
	"errors"
	"io"
	"log"
)

// StandardLogHandler send messages to a io.Writer using the standard logger.
type StandardLogHandler struct {
	stdlog *log.Logger
	output io.Writer
}

// NewStandardLogHandler creates a handler to send the logs to io.Writer through a standard logger.
func NewStandardLogHandler(output io.Writer, prefix string, flag int) *StandardLogHandler {
	handler := &StandardLogHandler{
		stdlog: log.New(output, prefix, flag),
		output: output,
	}
	return handler
}

// LogEntry sends a log entry with the specified level.
func (h *StandardLogHandler) LogEntry(logEntry LogEntry) error {
	return h.stdlog.Output(logEntry.Calldepth+2, logEntry.GetMessageWithLevelPrefix())
}

// SetOutput sets the output destination for the logger.
func (h *StandardLogHandler) SetOutput(output io.Writer) *StandardLogHandler {
	h.output = output
	h.stdlog.SetOutput(output)
	return h
}

// SetPrefix sets a prefix on every log message
func (h *StandardLogHandler) SetPrefix(prefix string) Handler {
	h.stdlog.SetPrefix(prefix)
	return h
}

// Close the output writer
func (h *StandardLogHandler) Close() error {
	closer, ok := h.output.(io.Closer)
	if ok {
		return closer.Close()
	}
	return errors.New("output cannot be closed")
}

// Verify interface
var (
	_ Handler = &StandardLogHandler{}
)
