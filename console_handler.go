package clog

import (
	"log"

	"github.com/fatih/color"
)

const (
	numLevels = 5 // Trace, Debug, Info, Warn, Error
)

// ConsoleHandler logs messages to the console (in colour)
type ConsoleHandler struct {
	colorMaps map[string][numLevels]*color.Color
	levelMap  [numLevels]*color.Color
	logger    *log.Logger
}

// NewConsoleHandler creates a new handler to send logs to the console
func NewConsoleHandler(prefix string, flag int) *ConsoleHandler {
	console := &ConsoleHandler{
		logger: log.New(color.Output, prefix, flag),
	}
	console.init()
	return console
}

func (h *ConsoleHandler) init() {
	h.colorMaps = map[string][numLevels]*color.Color{
		"none": {
			LevelTrace:   color.New(color.Reset),
			LevelDebug:   color.New(color.Reset),
			LevelInfo:    color.New(color.Reset),
			LevelWarning: color.New(color.Bold),
			LevelError:   color.New(color.Bold),
		},
		"light": {
			LevelTrace:   color.New(color.Faint),
			LevelDebug:   color.New(color.FgGreen),
			LevelInfo:    color.New(color.FgCyan),
			LevelWarning: color.New(color.FgMagenta, color.Bold),
			LevelError:   color.New(color.FgRed, color.Bold),
		},
		"dark": {
			LevelTrace:   color.New(color.Faint),
			LevelDebug:   color.New(color.FgHiGreen),
			LevelInfo:    color.New(color.FgHiCyan),
			LevelWarning: color.New(color.FgHiMagenta, color.Bold),
			LevelError:   color.New(color.FgHiRed, color.Bold),
		},
	}
	h.levelMap = h.colorMaps["light"]
}

// SetTheme sets a new color theme, and returns the ConsoleHandler for chaining.
// Accepted values are: "none", "light", "dark"
func (h *ConsoleHandler) SetTheme(theme string) *ConsoleHandler {
	var ok bool
	h.levelMap, ok = h.colorMaps[theme]
	if !ok {
		h.levelMap = h.colorMaps["none"]
	}
	return h
}

// Colouring activate of deactivate displaying messages in colour in the console, and returns the ConsoleHandler for chaining.
func (h *ConsoleHandler) Colouring(colouring bool) *ConsoleHandler {
	color.NoColor = !colouring
	return h
}

// SetPrefix sets a prefix on every log message, and returns the Handler interface for chaining.
func (h *ConsoleHandler) SetPrefix(prefix string) Handler {
	h.logger.SetPrefix(prefix)
	return h
}

// LogEntry sends a log entry with the specified level
func (h *ConsoleHandler) LogEntry(logEntry LogEntry) error {
	if h.levelMap[logEntry.Level] == nil {
		return ErrMessageDiscarded
	}
	return h.logger.Output(logEntry.Calldepth+2, h.levelMap[logEntry.Level].Sprint(logEntry.GetMessage()))
}

// Verify interface
var (
	_ Handler = &ConsoleHandler{}
)
