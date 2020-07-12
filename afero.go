package clog

import (
	"os"

	"github.com/spf13/afero"
)

// AferoHandler logs messages to a file
type AferoHandler struct {
	*StandardLogHandler
	file afero.File
}

// NewAferoHandler creates a new file logger through afero Fs.
//  Remember to Close() the logger at the end
func NewAferoHandler(fs afero.Fs, filename string, prefix string, flag int) (*AferoHandler, error) {
	file, err := fs.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	// standard output is managed by a StandardLogHandler
	handler := &AferoHandler{
		StandardLogHandler: NewStandardLogHandler(file, prefix, flag),
		file:               file,
	}
	return handler, nil
}

// Close the logfile when no longer needed
//  please note this method reinstate the standard console output as default
func (l *AferoHandler) Close() {
	if l.file != nil {
		l.file.Sync()
		l.file.Close()
		l.file = nil
	}
	// make sure any other call to the logger won't panic
	l.stdlog.SetOutput(os.Stdout)
}

// Verify interface
var (
	_ Handler = &AferoHandler{}
)
