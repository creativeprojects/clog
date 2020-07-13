package clog

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAferoLogExist(t *testing.T) {
	filename := "/var/log/test1.log"
	testFS := afero.NewMemMapFs()
	// create test directories
	testFS.MkdirAll("/var/log", 0755)

	handler, err := NewAferoHandler(testFS, filename, "", 0)
	require.NoError(t, err)

	logger := NewLogger(handler)

	logger.Log(LevelInfo, "one", "two", "three")
	logger.Debug("one", "two", "three")
	logger.Info("one", "two", "three")
	logger.Warning("one", "two", "three")
	logger.Error("one", "two", "three")

	logger.Logf(LevelInfo, "%d %d %d", 1, 2, 3)
	logger.Debugf("%d %d %d", 1, 2, 3)
	logger.Infof("%d %d %d", 1, 2, 3)
	logger.Warningf("%d %d %d", 1, 2, 3)
	logger.Errorf("%d %d %d", 1, 2, 3)

	handler.Close()
	if _, err := testFS.Stat(filename); err != nil || os.IsNotExist(err) {
		t.Errorf("cannot find log file %s", filename)
	}
}

func TestAferoDeleteLogFile(t *testing.T) {
	filename := "/var/log/test2.log"
	testFS := afero.NewMemMapFs()
	// create test directories
	testFS.MkdirAll("/var/log", 0755)

	handler, err := NewAferoHandler(testFS, filename, "", 0)
	require.NoError(t, err)

	logger := NewLogger(handler)

	logger.Log(LevelInfo, "one", "two", "three")

	// apparently the file can be deleted...
	err = testFS.Remove(filename)
	assert.NoError(t, err)

	// the logger should stay silent
	logger.Logf(LevelInfo, "%d %d %d", 1, 2, 3)
	// but the handler should return an error, and it doesn't, yet
	err = handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"test"},
	})
	assert.NoError(t, err)

	handler.Close()
	if _, err := testFS.Stat(filename); err == nil || os.IsExist(err) {
		t.Errorf("logfile still exists: %s", filename)
	}
}
