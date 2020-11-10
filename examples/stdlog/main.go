package main

import (
	"github.com/creativeprojects/clog"
)

//
// Example how to use a StandardLogger as a drop-in replacement for log.Logger
//

// outputLogger is an interface used to bring you own logger
type outputLogger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicln(v ...interface{})
	Panicf(format string, v ...interface{})
}

// someLibrary could be a library that would send logs to an interface where you can plug-in a standard logger
type someLibrary struct {
	Logger outputLogger
}

func (s *someLibrary) doStuff() {
	s.Logger.Print("starting doing stuff")
	s.Logger.Println("keep doing stuff")
	s.Logger.Printf("finished %s %s", "doing", "stuff")
}

func main() {
	handler := clog.NewConsoleHandler("library: ", 0)
	lib := &someLibrary{
		Logger: clog.NewStandardLogger(clog.LevelInfo, handler),
	}

	lib.doStuff()
}
