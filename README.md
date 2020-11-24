# golem
A leveled logger in json format built with golang.

# Features
- Simple
- Easy
- Leveled logger
- Json format
- Log threshold

# Installation
`go get -u github.com/bmf-san/golem`

# Example
```go
package main

import (
	"fmt"
	"time"

	golem "github.com/bmf-san/golem"
)

func main() {
	fmt.Println("Info Level")
	logger := golem.NewLogger(golem.LevelInfo, time.FixedZone("Asia/Tokyo", 9*60*60))
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")

	fmt.Println("Warn Level")
	logger = golem.NewLogger(golem.LevelWarn, time.FixedZone("Asia/Tokyo", 9*60*60))
	logger.Info("info") // This doesn't output
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")

	fmt.Println("Error Level")
	logger = golem.NewLogger(golem.LevelError, time.FixedZone("Asia/Tokyo", 9*60*60))
	logger.Info("info") // This doesn't output
	logger.Warn("warn") // This doesn't output
	logger.Error("error")
	logger.Fatal("fatal")

	fmt.Println("Fatal Level")
	logger = golem.NewLogger(golem.LevelFatal, time.FixedZone("Asia/Tokyo", 9*60*60))
	logger.Info("info")   // This doesn't output
	logger.Warn("warn")   // This doesn't output
	logger.Error("error") // This doesn't output
	logger.Fatal("fatal")
}

```

# Example outputs
```
{"level":"fatal","time":"2020-05-25T13:59:40.191444Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:10","message":"fatal"}
{"level":"error","time":"2020-05-25T13:59:40.191732Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:11","message":"error"}
{"level":"warn","time":"2020-05-25T13:59:40.191747Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:12","message":"warn"}
{"level":"info","time":"2020-05-25T13:59:40.191756Z","file":"/Users/bmf/localdev/godev/src/github.com/bmf-san/golem/_examples/main.go:13","message":"info"}
```