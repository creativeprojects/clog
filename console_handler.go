package clog

import (
	"log"
	"os"

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
		logger: log.New(os.Stdout, prefix, flag),
	}
	console.init()
	return console
}

func (l *ConsoleHandler) init() {
	l.colorMaps = map[string][numLevels]*color.Color{
		"none": {
			LevelTrace:   nil,
			LevelDebug:   nil,
			LevelInfo:    nil,
			LevelWarning: color.New(color.Bold),
			LevelError:   color.New(color.Bold),
		},
		"light": {
			LevelTrace:   nil,
			LevelDebug:   color.New(color.FgGreen),
			LevelInfo:    color.New(color.FgCyan),
			LevelWarning: color.New(color.FgMagenta, color.Bold),
			LevelError:   color.New(color.FgRed, color.Bold),
		},
		"dark": {
			LevelTrace:   nil,
			LevelDebug:   color.New(color.FgHiGreen),
			LevelInfo:    color.New(color.FgHiCyan),
			LevelWarning: color.New(color.FgHiMagenta, color.Bold),
			LevelError:   color.New(color.FgHiRed, color.Bold),
		},
	}
	l.levelMap = l.colorMaps["light"]
}

// SetTheme sets the dark or light theme
func (l *ConsoleHandler) SetTheme(theme string) {
	var ok bool
	l.levelMap, ok = l.colorMaps[theme]
	if !ok {
		l.levelMap = l.colorMaps["none"]
	}
}

// Colouring activate of deactivate displaying messages in colour in the console
func (l *ConsoleHandler) Colouring(colouring bool) {
	color.NoColor = !colouring
}

// SetPrefix sets a prefix on every log message
func (l *ConsoleHandler) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}

// LogEntry sends a log entry with the specified level
func (l *ConsoleHandler) LogEntry(logEntry LogEntry) error {
	l.setColor(l.levelMap[logEntry.Level])
	defer l.unsetColor()
	return l.logger.Output(logEntry.Calldepth+2, logEntry.GetMessage())
}

func (l *ConsoleHandler) setColor(c *color.Color) {
	if c != nil {
		c.Set()
	}
}

func (l *ConsoleHandler) unsetColor() {
	color.Unset()
}

// Verify interface
var (
	_ Handler = &ConsoleHandler{}
)
