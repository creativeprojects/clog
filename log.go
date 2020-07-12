package clog

var (
	defaultLogger = NewLogger(NewConsoleHandler())
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
	defaultLogger.Log(level, v...)
}

// Logf sends a log entry with the specified level
func Logf(level LogLevel, format string, v ...interface{}) {
	defaultLogger.Logf(level, format, v...)
}

// Debug sends debugging information
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Debugf sends debugging information
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Info logs some noticeable information
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Infof logs some noticeable information
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Warning send some important message to the console
func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

// Warningf send some important message to the console
func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

// Error sends error information to the console
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Errorf sends error information to the console
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}
