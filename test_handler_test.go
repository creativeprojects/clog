package clog

import "testing"

func TestTestHandler(t *testing.T) {
	SetTestLog(t)
	defer CloseTestLog()

	Log(LevelInfo, "one", "two", "three")
	Debug("one", "two", "three")
	Info("one", "two", "three")
	Warning("one", "two", "three")
	Error("one", "two", "three")

	Logf(LevelInfo, "%d %d %d", 1, 2, 3)
	Debugf("%d %d %d", 1, 2, 3)
	Infof("%d %d %d", 1, 2, 3)
	Warningf("%d %d %d", 1, 2, 3)
	Errorf("%d %d %d", 1, 2, 3)
}
