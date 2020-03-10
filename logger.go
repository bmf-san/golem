package golem

import (
	"encoding/json"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Level is log levels.
type Level int8

const (
	// FatalLevel is an abend error.
	FatalLevel Level = iota
	// ErrorLevel is unexpected runtime error.
	ErrorLevel
	// WarnLevel is warning.
	WarnLevel
	// InfoLevel is something notable infomation.
	InfoLevel
)

// String converts constant to string.
func (l Level) String() string {
	switch l {
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case InfoLevel:
		return "info"
	}
	return ""
}

// A Logger represents a logger.
type Logger struct {
	mu    sync.Mutex
	out   io.Writer
	entry *Entry
}

// An Entry represents a entry.
type Entry struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"` // UTC
	File    string    `json:"file"`
	Message string    `json:"message"`
}

// NewLogger creates a logger.
func NewLogger() *Logger {
	return &Logger{
		entry: &Entry{},
	}
}

// Fatal outputs a fatal level log.
func (l *Logger) Fatal(message string) {
	l.SetOutput(os.Stderr)
	l.OutputJSON(FatalLevel.String(), message)
}

// Error outputs a error level log.
func (l *Logger) Error(message string) {
	l.SetOutput(os.Stderr)
	l.OutputJSON(ErrorLevel.String(), message)
}

// Warn outputs a warn level log.
func (l *Logger) Warn(message string) {
	l.SetOutput(os.Stdout)
	l.OutputJSON(WarnLevel.String(), message)
}

// Info outputs a info level log.
func (l *Logger) Info(message string) {
	l.SetOutput(os.Stdout)
	l.OutputJSON(InfoLevel.String(), message)
}

// SetOutput sets the output.
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

// formatHeader writes log header.
func (l *Logger) formatHeader(t time.Time, file string, line int) {
	l.entry.Time = t.UTC()
	l.entry.File = file + ":" + strconv.Itoa(line)
}

// OutputJSON outputs logs.
func (l *Logger) OutputJSON(level string, msg string) error {
	l.entry.Level = level
	l.entry.Message = msg

	now := time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()

	l.mu.Unlock()
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	l.mu.Lock()

	l.formatHeader(now, file, line)

	bytes, err := json.Marshal(l.entry)
	if err != nil {
		return err
	}
	bytes = append(bytes, '\n')

	_, err = l.out.Write(bytes)
	return err
}
