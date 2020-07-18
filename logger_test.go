package golem

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	cases := []struct {
		actual   string
		expected string
	}{
		{
			actual:   FatalLevel.String(),
			expected: "fatal",
		},
		{
			actual:   ErrorLevel.String(),
			expected: "error",
		},
		{
			actual:   WarnLevel.String(),
			expected: "warn",
		},
		{
			actual:   InfoLevel.String(),
			expected: "info",
		},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("actual: %v expected: %v\n", c.actual, c.expected)
		}
	}
}

func TestNewLogger(t *testing.T) {
	actual := NewLogger()
	expected := &Logger{
		entry: &Entry{},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %v expected: %v\n", actual, expected)
	}
}

func TestAll(t *testing.T) {
	logger := NewLogger()

	logger.Fatal("fatal")
	logger.Error("error")
	logger.Warn("warn")
	logger.Info("info")
}
