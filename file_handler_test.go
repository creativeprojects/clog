package clog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileLogExist(t *testing.T) {
	filename := "test1.log"
	defer os.Remove(filename)

	handler, err := NewFileHandler(filename, "", 0)
	require.NoError(t, err)

	logger := NewLogger(handler)

	logger.Log(LevelInfo, "one", "two", "three")
	logger.Logf(LevelInfo, "%d %d %d", 1, 2, 3)

	handler.Close()
	if _, err := os.Stat(filename); err != nil || os.IsNotExist(err) {
		t.Errorf("cannot find log file %s", filename)
	}
}

func TestFileCloseLogFile(t *testing.T) {
	filename := "test2.log"

	handler, err := NewFileHandler(filename, "", 0)
	require.NoError(t, err)
	defer handler.Close()
	defer os.Remove(filename)

	logger := NewLogger(handler)

	logger.Log(LevelInfo, "one", "two", "three")

	// drastically close the file
	handler.file.Close()

	// the logger should stay silent
	logger.Logf(LevelInfo, "%d %d %d", 1, 2, 3)
	// but the handler should return an error
	err = handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"test"},
	})
	assert.Error(t, err)
}

func TestCloseFileHandler(t *testing.T) {
	filename := "test3.log"

	handler, err := NewFileHandler(filename, "", 0)
	require.NoError(t, err)
	defer handler.Close()
	defer os.Remove(filename)

	logger := NewLogger(handler)

	logger.Log(LevelInfo, "one", "two", "three")

	// close the handler properly
	handler.Close()

	// but the handler should not return an error as the outpout should have been diverted
	err = handler.LogEntry(LogEntry{
		Level:  LevelDebug,
		Values: []interface{}{"test"},
	})
	assert.NoError(t, err)
}
