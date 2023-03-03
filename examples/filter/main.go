package main

import (
	"fmt"
	"github.com/creativeprojects/clog"
)

func main() {
	log := clog.NewFilteredConsoleLogger(clog.LevelInfo)

	log.Info("will be displayed")
	log.Debug("will be discarded")
	log.Trace("will be discarded")
	log.Trace(func() string { return "will not be called" })

	log.Info(fmt.Sprintf, "generated and displayed(%d)", 1)
	log.Infof("generated and displayed(%d)", 2)
}
