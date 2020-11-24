package golem

import (
	"reflect"
	"testing"
	"time"
)

func TestLevelText(t *testing.T) {
	cases := []struct {
		expected string
		item     int
	}{
		{
			expected: LevelTextInfo,
			item:     0,
		},
		{
			expected: LevelTextWarn,
			item:     1,
		},
		{
			expected: LevelTextError,
			item:     2,
		},
		{
			expected: LevelTextFatal,
			item:     3,
		},
	}

	for _, c := range cases {
		actual := LevelText(c.item)
		if actual != c.expected {
			t.Errorf("actual: %v expected: %v\n", actual, c.expected)
		}
	}
}

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
