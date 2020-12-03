package main

import (
	"log"
	"sync"

	"github.com/creativeprojects/clog"
)

func main() {
	wg := &sync.WaitGroup{}
	logger := clog.NewLogger(clog.NewConsoleHandler("", log.Lshortfile|log.LstdFlags))
	for round := 0; round < 10; round++ {
		for i := clog.LevelTrace; i <= clog.LevelError; i++ {
			wg.Add(1)
			go func(level clog.LogLevel) {
				logger.Log(level, "Test message level ", level.String())
				wg.Done()
			}(clog.LogLevel(i))
		}
	}
	wg.Wait()
}
