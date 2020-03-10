# golem
A leveled logger in json format built with golang.

# Installation
`go get -u github.com/bmf-san/golem`

# Example
```go
package main

import (
	golem "github.com/bmf-san/golem"
)

func main() {
	logger := golem.NewLogger()

	logger.Fatal("fatal")
	logger.Error("error")
	logger.Warn("warn")
	logger.Info("info")
}
```

```
{"level":"fatal","time":"2020-05-25T13:59:40.191444Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:10","message":"fatal"}
{"level":"error","time":"2020-05-25T13:59:40.191732Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:11","message":"error"}
{"level":"warn","time":"2020-05-25T13:59:40.191747Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:12","message":"warn"}
{"level":"info","time":"2020-05-25T13:59:40.191756Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:13","message":"info"}
```