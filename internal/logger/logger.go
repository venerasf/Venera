package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelPanic:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Module    string                 `json:"module"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// Logger provides structured logging capabilities
type Logger struct {
	mu       sync.Mutex
	logPath  string
	minLevel LogLevel
	file     *os.File
}

var defaultLogger *Logger

// InitLogger initializes the default logger
func InitLogger(logPath string, minLevel LogLevel) error {
	// Ensure directory exists
	dir := filepath.Dir(logPath)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	defaultLogger = &Logger{
		logPath:  logPath,
		minLevel: minLevel,
		file:     f,
	}

	return nil
}

// Close closes the logger file
func Close() error {
	if defaultLogger != nil && defaultLogger.file != nil {
		return defaultLogger.file.Close()
	}
	return nil
}

// log writes a log entry
func (l *Logger) log(level LogLevel, module, message string, fields map[string]interface{}) {
	if level < l.minLevel {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level.String(),
		Module:    module,
		Message:   sanitizeMessage(message),
		Fields:    sanitizeFields(fields),
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Write JSON to file if available
	if l.file != nil {
		data, err := json.Marshal(entry)
		if err == nil {
			l.file.Write(data)
			l.file.Write([]byte("\n"))
		}
	}
}

// Debug logs a debug message
func Debug(module, message string, fields ...map[string]interface{}) {
	if defaultLogger != nil {
		var f map[string]interface{}
		if len(fields) > 0 {
			f = fields[0]
		}
		defaultLogger.log(LevelDebug, module, message, f)
	}
}

// Info logs an info message
func Info(module, message string, fields ...map[string]interface{}) {
	if defaultLogger != nil {
		var f map[string]interface{}
		if len(fields) > 0 {
			f = fields[0]
		}
		defaultLogger.log(LevelInfo, module, message, f)
	}
}

// Warn logs a warning message
func Warn(module, message string, fields ...map[string]interface{}) {
	if defaultLogger != nil {
		var f map[string]interface{}
		if len(fields) > 0 {
			f = fields[0]
		}
		defaultLogger.log(LevelWarn, module, message, f)
	}
}

// Error logs an error message
func Error(module, message string, fields ...map[string]interface{}) {
	if defaultLogger != nil {
		var f map[string]interface{}
		if len(fields) > 0 {
			f = fields[0]
		}
		defaultLogger.log(LevelError, module, message, f)
	}
}

// Panic logs a panic message
func Panic(module, message string, fields ...map[string]interface{}) {
	if defaultLogger != nil {
		var f map[string]interface{}
		if len(fields) > 0 {
			f = fields[0]
		}
		defaultLogger.log(LevelPanic, module, message, f)
	}
}

// sanitizeMessage removes or masks sensitive information from messages
func sanitizeMessage(msg string) string {
	// Remove potential passwords, tokens, keys
	msg = strings.ReplaceAll(msg, "password=", "password=***")
	msg = strings.ReplaceAll(msg, "token=", "token=***")
	msg = strings.ReplaceAll(msg, "api_key=", "api_key=***")
	msg = strings.ReplaceAll(msg, "secret=", "secret=***")
	return msg
}

// sanitizeFields removes sensitive keys from log fields
func sanitizeFields(fields map[string]interface{}) map[string]interface{} {
	if fields == nil {
		return nil
	}

	sanitized := make(map[string]interface{})
	sensitiveKeys := map[string]bool{
		"password":   true,
		"token":      true,
		"api_key":    true,
		"secret":     true,
		"apikey":     true,
		"auth":       true,
		"auth_key":   true,
		"priv_key":   true,
		"privatekey": true,
	}

	for k, v := range fields {
		lowerKey := strings.ToLower(k)
		if sensitiveKeys[lowerKey] {
			sanitized[k] = "***"
		} else {
			sanitized[k] = v
		}
	}

	return sanitized
}

// LogWriter returns an io.Writer for the logger
func LogWriter() io.Writer {
	if defaultLogger != nil && defaultLogger.file != nil {
		return defaultLogger.file
	}
	return os.Stderr
}
