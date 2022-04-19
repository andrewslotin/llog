package llog

import (
	"bytes"
	"io"
	"sync"
)

// Level is the log level used by the Writer.
type Level int8

// Valid log levels for the Writer
const (
	FatalLevel Level = iota - 1
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// String returns the string representation of the log level.
func (l Level) String() string {
	switch l {
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

// Writer wraps an io.Writer filtering out log messages based on the current log level.
type Writer struct {
	io.Writer

	mu  sync.RWMutex
	lvl Level
}

// NewWriter initializes a new Writer with the given io.Writer and log level.
func NewWriter(out io.Writer, lvl Level) *Writer {
	return &Writer{
		lvl:    lvl,
		Writer: out,
	}
}

// Write determines the log level of the message and writes it to the underlying io.Writer if the log level is greater than or equal to the Writer's log level.
// The log level of the message is determined by the case-sensitive message prefix:
// * messages starting with "missing" or "fatal" are considered fatal errors
// * messages starting with "error" or "failed" are considered errors
// * messages starting with "warn" are considered warnings
// * messages starting with "debug:" (including the colon) are considered debug messages
// * anything else is considered to be an info
func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.RLock()
	lvl := w.lvl
	w.mu.RUnlock()

	if determineLevel(p) > lvl {
		return 0, nil
	}

	return w.Writer.Write(p)
}

// SetLevel sets the log level of the Writer.
func (w *Writer) SetLevel(lvl Level) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.lvl = lvl
}

func determineLevel(msg []byte) Level {
	switch {
	case bytes.HasPrefix(msg, []byte("fatal")) || bytes.HasPrefix(msg, []byte("missing")):
		return FatalLevel
	case bytes.HasPrefix(msg, []byte("error")) || bytes.HasPrefix(msg, []byte("failed")):
		return ErrorLevel
	case bytes.HasPrefix(msg, []byte("warn")):
		return WarnLevel
	case bytes.HasPrefix(msg, []byte("debug:")):
		return DebugLevel
	default:
		return InfoLevel
	}
}
