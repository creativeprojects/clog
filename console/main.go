package main

import (
	"fmt"
	"log"

	"github.com/creativeprojects/clog"
)

func main() {
	fmt.Println("Test quick access methods:")
	clog.Debug("debug ", "message")
	clog.Debugf("%s %s", "another", "one")
	clog.Info("info ", "message")
	clog.Infof("%s %s", "another", "one")
	clog.Warning("warning ", "message")
	clog.Warningf("%s %s", "another", "one")
	clog.Error("error ", "message")
	clog.Errorf("%s %s", "another", "one")

	fmt.Println("Test creating a new console logger")
	logger := clog.NewLogger(clog.NewConsoleHandler("", log.Lshortfile))
	for i := clog.LevelDebug; i <= clog.LevelError; i++ {
		level := clog.LogLevel(i)
		logger.Log(level, "Test message level ", level.String())
	}
}
