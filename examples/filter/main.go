package main

import "github.com/creativeprojects/clog"

func main() {
	log := clog.NewFilteredConsoleLogger(clog.LevelInfo)

	log.Debug("will be discarded")
	log.Info("will be displayed")
}
