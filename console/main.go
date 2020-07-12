package main

import "github.com/creativeprojects/clog"

func main() {
	clog.Debug("debug", "message")
	clog.Debugf("%s %s", "another", "one")
	clog.Info("info", "message")
	clog.Infof("%s %s", "another", "one")
	clog.Warning("warning", "message")
	clog.Warningf("%s %s", "another", "one")
	clog.Error("error", "message")
	clog.Errorf("%s %s", "another", "one")
}
