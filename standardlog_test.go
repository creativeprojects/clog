package clog

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardLogPrint(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewStandardLogger(LevelInfo, handler)
	logger.Print("one", "message")
	assert.ElementsMatch(t, []string{"onemessage"}, handler.log)
}

func TestStandardLogPrintln(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewStandardLogger(LevelInfo, handler)
	logger.Println("one", "message")
	assert.ElementsMatch(t, []string{"onemessage"}, handler.log)
}

func TestStandardLogPrintf(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewStandardLogger(LevelInfo, handler)
	logger.Printf("%s %s", "one", "message")
	assert.ElementsMatch(t, []string{"one message"}, handler.log)
}

func TestStandardLogFatal(t *testing.T) {
	// Nice trick found here: https://stackoverflow.com/a/33404435
	if os.Getenv("BE_CRASHER") == "1" {
		NewStandardLogger(LevelError, &DiscardHandler{}).Fatal("")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestStandardLogFatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestStandardLogFatalln(t *testing.T) {
	// Nice trick found here: https://stackoverflow.com/a/33404435
	if os.Getenv("BE_CRASHER") == "1" {
		NewStandardLogger(LevelError, &DiscardHandler{}).Fatalln("")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestStandardLogFatalln")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestStandardLogFatalf(t *testing.T) {
	// Nice trick found here: https://stackoverflow.com/a/33404435
	if os.Getenv("BE_CRASHER") == "1" {
		NewStandardLogger(LevelError, &DiscardHandler{}).Fatalf("")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestStandardLogFatalf")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestStandardLogPanic(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewStandardLogger(LevelInfo, handler)
	assert.Panics(t, func() {
		logger.Panic("one", "message")
	})
	assert.ElementsMatch(t, []string{"onemessage"}, handler.log)
}

func TestStandardLogPanicln(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewStandardLogger(LevelInfo, handler)
	assert.Panics(t, func() {
		logger.Panicln("one", "message")
	})
	assert.ElementsMatch(t, []string{"onemessage"}, handler.log)
}

func TestStandardLogPanicf(t *testing.T) {
	handler := NewMemoryHandler()
	logger := NewStandardLogger(LevelInfo, handler)
	assert.Panics(t, func() {
		logger.Panicf("%s %s", "one", "message")
	})
	assert.ElementsMatch(t, []string{"one message"}, handler.log)
}
