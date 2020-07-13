package clog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilteredLogger(t *testing.T) {
	expected := []string{
		"DEBUG 0 >= 0",
		"INFO  1 >= 0",
		"WARN  2 >= 0",
		"ERROR 3 >= 0",
		"INFO  1 >= 1",
		"WARN  2 >= 1",
		"ERROR 3 >= 1",
		"WARN  2 >= 2",
		"ERROR 3 >= 2",
		"ERROR 3 >= 3",
	}

	buffer := &bytes.Buffer{}
	handler := NewStandardLogHandler(buffer, "", 0)

	for minLevel := LevelDebug; minLevel <= LevelError; minLevel++ {
		logger := NewLogger(NewLevelFilter(minLevel, handler))
		for logLevel := LevelDebug; logLevel <= LevelError; logLevel++ {
			logger.Logf(logLevel, "%d >= %d", logLevel, minLevel)
		}
	}
	logs := []string{}
	for line, err := buffer.ReadString('\n'); err == nil; line, err = buffer.ReadString('\n') {
		logs = append(logs, strings.Trim(line, "\n"))
	}
	assert.ElementsMatch(t, expected, logs)
}
