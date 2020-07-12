package clog

import (
	"log"
	"os"

	"github.com/fatih/color"
)

const (
	numLevels = 4 // Debug, Info, Warn, Error
)

// ConsoleHandler logs messages to the console (in colour)
type ConsoleHandler struct {
	colorMaps map[string][numLevels]*color.Color
	levelMap  [numLevels]*color.Color
	logger    *log.Logger
}

// NewConsoleHandler creates a new handler to send logs to the console
func NewConsoleHandler() *ConsoleHandler {
	console := &ConsoleHandler{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
	console.init()
	return console
}

func (l *ConsoleHandler) init() {
	l.colorMaps = map[string][numLevels]*color.Color{
		"none": {
			LevelDebug:   nil,
			LevelInfo:    nil,
			LevelWarning: color.New(color.Bold),
			LevelError:   color.New(color.Bold),
		},
		"light": {
			LevelDebug:   color.New(color.FgGreen),
			LevelInfo:    color.New(color.FgCyan),
			LevelWarning: color.New(color.FgMagenta, color.Bold),
			LevelError:   color.New(color.FgRed, color.Bold),
		},
		"dark": {
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

// Colorize activate of deactivate colouring
func (l *ConsoleHandler) Colorize(colorize bool) {
	color.NoColor = !colorize
}

// Log sends a log entry with the specified level
func (l *ConsoleHandler) Log(logEntry LogEntry) error {
	if logEntry.Format == "" {
		l.message(l.levelMap[logEntry.Level], logEntry.Values...)
		return nil
	}
	l.messagef(l.levelMap[logEntry.Level], logEntry.Format, logEntry.Values...)
	return nil
}

func (l *ConsoleHandler) message(c *color.Color, v ...interface{}) {
	l.setColor(c)
	l.logger.Println(v...)
	l.unsetColor()
}

func (l *ConsoleHandler) messagef(c *color.Color, format string, v ...interface{}) {
	l.setColor(c)
	l.logger.Printf(format+"\n", v...)
	l.unsetColor()
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
