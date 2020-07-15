package clog

import (
	"os"
)

// FileHandler logs messages to a file
type FileHandler struct {
	*StandardLogHandler
	file *os.File
}

// NewFileHandler creates a new file logger
//  Remember to Close() the logger at the end
func NewFileHandler(filename string, prefix string, flag int) (*FileHandler, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	// standard output is managed by a StandardLogHandler
	handler := &FileHandler{
		StandardLogHandler: NewStandardLogHandler(file, prefix, flag),
		file:               file,
	}
	return handler, nil
}

// Close the logfile when no longer needed
//  please note this method reinstate the standard console output as default
func (l *FileHandler) Close() {
	if l.file != nil {
		l.file.Sync()
		l.file.Close()
		l.file = nil
	}
	// make sure any other call to the handler won't panic
	l.SetOutput(os.Stderr)
}

// Verify interface
var (
	_ Handler = &FileHandler{}
)
