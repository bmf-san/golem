package golem

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

const (
	// LevelInfo is something notable infomation.
	LevelInfo int = iota
	// LevelWarn is warning.
	LevelWarn
	// LevelError is unexpected runtime error.
	LevelError
	// LevelFatal is an abend error.
	LevelFatal
)

const (
	// LevelTextInfo is the text for info.
	LevelTextInfo = "info"
	// LevelTextWarn is the text for warn.
	LevelTextWarn = "warn"
	// LevelTextError is the text for error.
	LevelTextError = "error"
	// LevelTextFatal is the text for fatal.
	LevelTextFatal = "fatal"
)

var levelText = map[int]string{
	LevelInfo:  LevelTextInfo,
	LevelWarn:  LevelTextWarn,
	LevelError: LevelTextError,
	LevelFatal: LevelTextFatal,
}

// LevelText returns the text of level.
func LevelText(level int) string {
	return levelText[level]
}

// A Logger represents a logger.
type Logger struct {
	mu        sync.Mutex
	out       io.Writer
	entry     *Entry
	threshold int
	location  *time.Location
}

// An Entry represents a entry.
type Entry struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"` // UTC
	Message string    `json:"message"`
}

// NewLogger creates a logger.
func NewLogger(threshold int, location *time.Location) *Logger {
	return &Logger{
		entry:     &Entry{},
		threshold: threshold,
		location:  location,
	}
}

// Info outputs a info level log.
func (l *Logger) Info(message string) {
	if LevelInfo >= l.threshold {
		l.SetOutput(os.Stdout)
		l.OutputJSON(LevelText(LevelInfo), message)
	}
}

// Warn outputs a warn level log.
func (l *Logger) Warn(message string) {
	if LevelWarn >= l.threshold {
		l.SetOutput(os.Stdout)
		l.OutputJSON(LevelText(LevelWarn), message)
	}
}

// Error outputs a error level log.
func (l *Logger) Error(message string) {
	if LevelError >= l.threshold {
		l.SetOutput(os.Stderr)
		l.OutputJSON(LevelText(LevelError), message)
	}
}

// Fatal outputs a fatal level log.
func (l *Logger) Fatal(message string) {
	if LevelFatal >= l.threshold {
		l.SetOutput(os.Stderr)
		l.OutputJSON(LevelText(LevelFatal), message)
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
