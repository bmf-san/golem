package golem

import (
	"reflect"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	threshold := LevelInfo
	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	actual := NewLogger(threshold, location)
	expected := &Logger{
		entry:     &Entry{},
		threshold: threshold,
		location:  location,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %v expected: %v\n", actual, expected)
	}
}

func TestAll(t *testing.T) {
	logger := NewLogger(LevelInfo, time.FixedZone("Asia/Tokyo", 9*60*60))

	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")
}
