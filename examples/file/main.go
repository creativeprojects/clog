package main

import (
	"log"

	"github.com/creativeprojects/clog"
)

func main() {
	handler, _ := clog.NewFileHandler("log.txt", "", log.Lshortfile|log.LstdFlags)
	// Close() will safely close the log file at the end
	defer handler.Close()

	logger := clog.NewLogger(handler)
	// you can set a new default logger to use the package methods
	clog.SetDefaultLogger(logger)

	clog.Debug("some debug message")
}
