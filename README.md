[![Godoc reference](https://godoc.org/github.com/creativeprojects/clog?status.svg)](https://pkg.go.dev/github.com/creativeprojects/clog)
![Build](https://github.com/creativeprojects/clog/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/creativeprojects/clog)](https://goreportcard.com/report/github.com/creativeprojects/clog)
[![codecov](https://codecov.io/gh/creativeprojects/clog/branch/master/graph/badge.svg?token=N11UE47ESS)](https://codecov.io/gh/creativeprojects/clog)


# clog (console-log)
All the fashionable loggers for Go tend to focus on structured logging, and that's perfectly fine: until you simply need a logger for a console application...

So here's yet another logger for Go:
- unstructured logging
- console logging in colour
- file logging
- simple to use
- filter your logs from 5 levels of severity (Trace, Debug, Info, Warn, Error)
- redirect your logs to an io.Writer
- get logs coming from an io.Writer
- using the logger from the standard library under the hood
- extensible (via handlers and middleware)
- unit test coverage of more than 90%
- drop-in replacement for the standard library logger

Have a look at the [examples](https://github.com/creativeprojects/clog/tree/master/examples) if you like the look of it

Here's a very simple one:

```go
package main

import (
	"fmt"
	"github.com/creativeprojects/clog"
)

func main() {
	log := clog.NewFilteredConsoleLogger(clog.LevelInfo)

	log.Info("will be displayed")
	log.Debug("will be discarded")
	log.Trace("will be discarded")
	log.Trace(func() string { return "will not be called" })

	log.Info(fmt.Sprintf, "generated and displayed(%d)", 1)
	log.Infof("generated and displayed(%d)", 2)
}

```

<img alt="example" src="https://github.com/creativeprojects/clog/raw/master/filter.png" width="300" title="FilteredHandler & ConsoleHandler">

Documentation available on [GoDoc](https://pkg.go.dev/github.com/creativeprojects/clog?tab=doc)
