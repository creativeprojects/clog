package clog

// Handler for a logger.
//
// The Log method should return an error if the handler didn't manage to save the log
// (file, remote, etc.)
// It's up to the parent handler to take action on the error: the default Logger is going to ignore it.
type Handler interface {
	Log(LogEntry) error
}
