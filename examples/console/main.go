package main

import (
	"fmt"
	"log"

	"github.com/creativeprojects/clog"
)

func main() {
	// you don't have to set a default logger to use the package functions,
	// only if you need to change the default options
	logger := clog.NewLogger(clog.NewConsoleHandler("", log.Lshortfile))
	clog.SetDefaultLogger(logger)

	fmt.Println("\n== Quick access using package functions ==")
	clog.Trace("trace ", "message")
	clog.Tracef("%s %s", "another", "trace")
	clog.Debug("debug ", "message")
	clog.Debugf("%s %s", "another", "debug")
	clog.Info("info ", "message")
	clog.Infof("%s %s", "another", "info")
	clog.Warning("warning ", "message")
	clog.Warningf("%s %s", "another", "warning")
	clog.Error("error ", "message")
	clog.Errorf("%s %s", "another", "error")
	clog.Log(clog.LevelDebug, "using Log function with Debug level")
	clog.Logf(clog.LevelError, "using Logf function with %s level", clog.LevelError)

	fmt.Println("\n== Using a new console logger ==")
	logger = clog.NewLogger(clog.NewConsoleHandler("", log.Lshortfile|log.LstdFlags))
	for i := clog.LevelTrace; i <= clog.LevelError; i++ {
		level := clog.LogLevel(i)
		logger.Log(level, "Test message level ", level.String())
	}

	fmt.Println("\n== Using a filtered console logger ==")
	// create a filtered handler and place it in between the existing logger and handler
	handler := clog.NewLevelFilter(clog.LevelWarning, logger.GetHandler())
	logger.SetHandler(handler)
	for i := clog.LevelDebug; i <= clog.LevelError; i++ {
		level := clog.LogLevel(i)
		logger.Log(level, "Test message level ", level.String())
	}
}
