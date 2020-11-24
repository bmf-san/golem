package golem

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

// Level is log levels.
type Level int8

const (
	// InfoLevel is something notable infomation.
	InfoLevel Level = iota
	// WarnLevel is warning.
	WarnLevel
	// ErrorLevel is unexpected runtime error.
	ErrorLevel
	// FatalLevel is an abend error.
	FatalLevel
)

// String converts constant to string.
func (l Level) String() string {
	switch l {
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	return ""
}

// A Logger represents a logger.
type Logger struct {
	mu        sync.Mutex
	out       io.Writer
	entry     *Entry
	threshold Level
	location  *time.Location
}

// An Entry represents a entry.
type Entry struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"` // UTC
	Message string    `json:"message"`
}

// NewLogger creates a logger.
func NewLogger(threshold Level, location *time.Location) *Logger {
	return &Logger{
		entry:     &Entry{},
		threshold: threshold,
		location:  location,
	}
}

// Info outputs a info level log.
func (l *Logger) Info(message string) {
	if InfoLevel >= l.threshold {
		l.SetOutput(os.Stdout)
		l.OutputJSON(InfoLevel.String(), message)
	}
}

// Warn outputs a warn level log.
func (l *Logger) Warn(message string) {
	if WarnLevel >= l.threshold {
		l.SetOutput(os.Stdout)
		l.OutputJSON(WarnLevel.String(), message)
	}
}

// Error outputs a error level log.
func (l *Logger) Error(message string) {
	if ErrorLevel >= l.threshold {
		l.SetOutput(os.Stderr)
		l.OutputJSON(ErrorLevel.String(), message)
	}
}

// Fatal outputs a fatal level log.
func (l *Logger) Fatal(message string) {
	if FatalLevel >= l.threshold {
		l.SetOutput(os.Stderr)
		l.OutputJSON(FatalLevel.String(), message)
	}
}

// SetOutput sets the output.
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

// OutputJSON outputs logs.
func (l *Logger) OutputJSON(level string, msg string) error {
	l.entry.Level = level
	l.entry.Message = msg
	l.entry.Time = time.Now().UTC().In(l.location)
	bytes, err := json.Marshal(l.entry)
	if err != nil {
		return err
	}
	bytes = append(bytes, '\n')

	_, err = l.out.Write(bytes)
	return err
}
