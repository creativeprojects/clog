package clog

// Handler for a logger.
//
// The LogEntry method should return an error if the handler didn't manage to save the log
// (file, remote, etc.)
// It's up to the parent handler to take action on the error: the default Logger is always going to ignore it.
type Handler interface {
	LogEntry(LogEntry) error
	SetPrefix(string)
}

// MiddlewareHandler is a Handler that act as a middleware => you can get and set the next handler in the chain
type MiddlewareHandler interface {
	GetHandler() Handler
	SetHandler(handler Handler)
}
