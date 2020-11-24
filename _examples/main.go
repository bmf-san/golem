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
