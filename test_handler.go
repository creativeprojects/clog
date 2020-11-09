package clog

// TestHandler redirects all the logs to the testing framework logger
type TestHandler struct {
	t TestLogInterface
}

// TestLogInterface for use with testing.B or testing.T
type TestLogInterface interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

// NewTestHandler instantiates a new logger redirecting to the test framework logger
// or any other implementation of TestLogInterface for that matter
func NewTestHandler(t TestLogInterface) *TestHandler {
	return &TestHandler{
		t: t,
	}
}

// SetTestLog install a test logger as the default logger.
// A test logger redirects all logs sent through the package methods to the Log/Logf methods of your test
//
// IMPORTANT: don't forget to CloseTestLog() at the end of the test
//
// Example:
//
// 	func TestLog(t *testing.T) {
// 		SetTestLog(t)
// 		defer CloseTestLog()
// 		// These two calls are equivalent:
// 		clog.Debug("debug message")
// 		t.Log("debug message")
// 	}
func SetTestLog(t TestLogInterface) {
	SetDefaultLogger(NewLogger(NewTestHandler(t)))
}

// CloseTestLog at the end of the test otherwise the logger will keep a reference on t.
// For a description on how to use it, see SetTestLog()
func CloseTestLog() {
	SetDefaultLogger(NewLogger(NewDiscardHandler()))
}

// LogEntry sends a log entry with the specified level
func (l *TestHandler) LogEntry(logEntry LogEntry) error {
	if logEntry.Format == "" {
		l.t.Log(append([]interface{}{logEntry.Level.String()}, logEntry.Values...)...)
		return nil
	}
	l.t.Logf(logEntry.Level.String()+" "+logEntry.Format, logEntry.Values...)
	return nil
}

// SetPrefix does nothing on the test handler
func (l *TestHandler) SetPrefix(prefix string) {}

// Verify interface
var (
	_ Handler = &TestHandler{}
)
