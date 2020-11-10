package clog

import "os"

// StandardLogger can be used when you need to plug-in a standard library logger (via an interface)
type StandardLogger struct {
	level    LogLevel
	handler  Handler
	exitFunc func()
}

// NewStandardLogger creates a new logger that can be used in place of a standard library logger (via an interface)
func NewStandardLogger(level LogLevel, handler Handler) *StandardLogger {
	return &StandardLogger{
		level:   level,
		handler: handler,
		exitFunc: func() {
			os.Exit(1)
		},
	}
}

// RegisterExitFunc allows using a different "exit" function when calling Fatal, Fatalln or Fatalf.
// If not specified, os.Exit(1) is used
func (l *StandardLogger) RegisterExitFunc(exitFunc func()) {
	if exitFunc != nil {
		l.exitFunc = exitFunc
	}
}

// Print writes the output for a logging event.
// Arguments are handled in the manner of fmt.Print.
// A newline is appended if the last character of s is not already a newline.
func (l *StandardLogger) Print(v ...interface{}) {
	l.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Values:    v,
	})
}

// Println writes the output for a logging event.
// Arguments are handled in the manner of fmt.Println.
func (l *StandardLogger) Println(v ...interface{}) {
	l.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Values:    v,
	})
}

// Printf writes the output for a logging event.
// Arguments are handled in the manner of fmt.Printf.
// A newline is appended if the last character of s is not already a newline.
func (l *StandardLogger) Printf(format string, v ...interface{}) {
	l.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Format:    format,
		Values:    v,
	})
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
// You can change the exit function with RegisterExitFunc if needed.
func (l *StandardLogger) Fatal(v ...interface{}) {
	l.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Values:    v,
	})
	l.exitFunc()
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
// You can change the exit function with RegisterExitFunc if needed.
func (l *StandardLogger) Fatalln(v ...interface{}) {
	l.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Values:    v,
	})
	l.exitFunc()
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
// You can change the exit function with RegisterExitFunc if needed.
func (l *StandardLogger) Fatalf(format string, v ...interface{}) {
	l.handler.LogEntry(LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Format:    format,
		Values:    v,
	})
	l.exitFunc()
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *StandardLogger) Panic(v ...interface{}) {
	entry := LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Values:    v,
	}
	l.handler.LogEntry(entry)
	panic(entry.GetMessage())
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *StandardLogger) Panicln(v ...interface{}) {
	entry := LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Values:    v,
	}
	l.handler.LogEntry(entry)
	panic(entry.GetMessage())
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *StandardLogger) Panicf(format string, v ...interface{}) {
	entry := LogEntry{
		Calldepth: 1,
		Level:     l.level,
		Format:    format,
		Values:    v,
	}
	l.handler.LogEntry(entry)
	panic(entry.GetMessage())
}
