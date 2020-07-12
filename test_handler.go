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
//  IMPORTANT: don't forget to run ClearTestLog() at the end of the test
func SetTestLog(t TestLogInterface) {
	SetDefaultLogger(NewLogger(NewTestHandler(t)))
}

// ClearTestLog at the end of the test otherwise the logger will keep a reference on t
func ClearTestLog() {
	SetDefaultLogger(NewLogger(&NullHandler{}))
}

// Log sends a log entry with the specified level
func (l *TestHandler) Log(logEntry LogEntry) error {
	if logEntry.Format == "" {
		l.t.Log(append([]interface{}{logEntry.Level.String()}, logEntry.Values...)...)
		return nil
	}
	l.t.Logf(logEntry.Level.String()+" "+logEntry.Format, logEntry.Values...)
	return nil
}

// Verify interface
var (
	_ Handler = &TestHandler{}
)
