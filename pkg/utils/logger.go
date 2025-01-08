package utils

import (
	"log"
	"os"
)

// Logger is a simple utility for logging with levels
type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Debug logs debug messages
func (l *Logger) Debug(message string, args ...interface{}) {
	l.debugLogger.Printf(message, args...)
}

// Info logs informational messages
func (l *Logger) Info(message string, args ...interface{}) {
	l.infoLogger.Printf(message, args...)
}

// Warn logs warning messages
func (l *Logger) Warn(message string, args ...interface{}) {
	l.warnLogger.Printf(message, args...)
}

// Error logs error messages
func (l *Logger) Error(message string, args ...interface{}) {
	l.errorLogger.Printf(message, args...)
}
