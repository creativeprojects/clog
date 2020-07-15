[![Build Status](https://travis-ci.com/creativeprojects/clog.svg?branch=master)](https://travis-ci.com/creativeprojects/clog)
[![Go Report Card](https://goreportcard.com/badge/github.com/creativeprojects/clog)](https://goreportcard.com/report/github.com/creativeprojects/clog)

# clog (console-log)
All the fashionable loggers for Go tend to focus on structured logging, and that's perfectly fine: until you simply need a logger for a console application...

So here's yet another logger for Go:
- unstructured logging
- console logging in colour
- simple to use
- using the logger from the standard library under the hood
- extensible (via handlers)

Have a look at the [examples](https://github.com/creativeprojects/clog/tree/master/examples) if you like the look of it

Here's a very simple one:

```go
import "github.com/creativeprojects/clog"

func main() {
	log := clog.NewFilteredConsoleLogger(clog.LevelInfo)

	log.Debug("will be discarded")
	log.Info("will be displayed")
}

```

![alt text](https://github.com/creativeprojects/clog/raw/master/filter.png "FilteredHandler & ConsoleHandler")

