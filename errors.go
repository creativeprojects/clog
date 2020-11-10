package clog

import "errors"

// errors
var (
	ErrNoRegisteredHandler = errors.New("no registered handler")
	ErrHandlerClosed       = errors.New("handler is closed")
	ErrMessageDiscarded    = errors.New("this message is not going anywhere")
	ErrNoPrimaryHandler    = errors.New("no primary handler registered")
	ErrNoBackupHandler     = errors.New("no backup handler registered")
)
