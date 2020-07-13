package clog

var (
	defaultLogger = NewConsoleLogger()
)

// SetDefaultLogger sets the logger used when using the package methods
func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

// GetDefaultLogger returns the logger used when using the package methods
func GetDefaultLogger() *Logger {
	return defaultLogger
}

// Log sends a log entry with the specified level
func Log(level LogLevel, v ...interface{}) {
	defaultLog(level, v...)
}

// Logf sends a log entry with the specified level
func Logf(level LogLevel, format string, v ...interface{}) {
	defaultLogf(level, format, v...)
}

// Debug sends debugging information
func Debug(v ...interface{}) {
	defaultLog(LevelDebug, v...)
}

// Debugf sends debugging information
func Debugf(format string, v ...interface{}) {
	defaultLogf(LevelDebug, format, v...)
}

// Info logs some noticeable information
func Info(v ...interface{}) {
	defaultLog(LevelInfo, v...)
}

// Infof logs some noticeable information
func Infof(format string, v ...interface{}) {
	defaultLogf(LevelInfo, format, v...)
}

// Warning send some important message to the console
func Warning(v ...interface{}) {
	defaultLog(LevelWarning, v...)
}

// Warningf send some important message to the console
func Warningf(format string, v ...interface{}) {
	defaultLogf(LevelWarning, format, v...)
}

// Error sends error information to the console
func Error(v ...interface{}) {
	defaultLog(LevelError, v...)
}

// Errorf sends error information to the console
func Errorf(format string, v ...interface{}) {
	defaultLogf(LevelError, format, v...)
}

// log is used to keep a constant calldepth
func defaultLog(level LogLevel, v ...interface{}) {
	defaultLogger.LogEntry(LogEntry{
		Calldepth: 2,
		Level:     level,
		Values:    v,
	})
}

// logf is used to keep a constant calldepth
func defaultLogf(level LogLevel, format string, v ...interface{}) {
	defaultLogger.LogEntry(LogEntry{
		Calldepth: 2,
		Level:     level,
		Format:    format,
		Values:    v,
	})
}
