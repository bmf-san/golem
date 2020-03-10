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
