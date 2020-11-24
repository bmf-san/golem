package golem

import (
	"reflect"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	cases := []struct {
		actual   string
		expected string
	}{
		{
			actual:   InfoLevel.String(),
			expected: "info",
		},
		{
			actual:   WarnLevel.String(),
			expected: "warn",
		},
		{
			actual:   ErrorLevel.String(),
			expected: "error",
		},
		{
			actual:   FatalLevel.String(),
			expected: "fatal",
		},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("actual: %v expected: %v\n", c.actual, c.expected)
		}
	}
}

func TestNewLogger(t *testing.T) {
	threshold := InfoLevel
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
	logger := NewLogger(InfoLevel, time.FixedZone("Asia/Tokyo", 9*60*60))

	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")
}
