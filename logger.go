package clog

// Logger frontend
type Logger struct {
	handler Handler
}

// NewLogger creates a new logger
func NewLogger(handler Handler) *Logger {
	return &Logger{
		handler: handler,
	}
}

// SetHandler sets a new handler for the logger
func (l *Logger) SetHandler(handler Handler) {
	l.handler = handler
}

// GetHandler returns the current handler used by the logger
func (l *Logger) GetHandler() Handler {
	return l.handler
}

// Log sends a log entry with the specified level
func (l *Logger) Log(level LogLevel, v ...interface{}) {
	l.handler.Log(LogEntry{
		Level:  level,
		Values: v,
	})
}

// Logf sends a log entry with the specified level
func (l *Logger) Logf(level LogLevel, format string, v ...interface{}) {
	l.handler.Log(LogEntry{
		Level:  level,
		Format: format,
		Values: v,
	})
}

// Debug sends debugging information
func (l *Logger) Debug(v ...interface{}) {
	l.Log(LevelDebug, v...)
}

// Debugf sends debugging information
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(LevelDebug, format, v...)
}

// Info logs some noticeable information
func (l *Logger) Info(v ...interface{}) {
	l.Log(LevelInfo, v...)
}

// Infof logs some noticeable information
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(LevelInfo, format, v...)
}

// Warning send some important message to the console
func (l *Logger) Warning(v ...interface{}) {
	l.Log(LevelWarning, v...)
}

// Warningf send some important message to the console
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Logf(LevelWarning, format, v...)
}

// Error sends error information to the console
func (l *Logger) Error(v ...interface{}) {
	l.Log(LevelError, v...)
}

// Errorf sends error information to the console
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(LevelError, format, v...)
}
