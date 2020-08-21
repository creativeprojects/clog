package clog

import (
	"io"
	"os"
)

var (
	stdout io.Writer
	stderr io.Writer
)

// Stdout returns a thread-safe io.Writer to os.Stdout.
// Calling this function multiple times always returns the same instance of io.Writer.
func Stdout() io.Writer {
	if stdout == nil {
		stdout = NewSafeWriter(os.Stdout)
	}
	return stdout
}

// Stderr returns a thread-safe io.Writer to os.Stderr.
// Calling this function multiple times always returns the same instance of io.Writer.
func Stderr() io.Writer {
	if stderr == nil {
		stderr = NewSafeWriter(os.Stderr)
	}
	return stderr
}
