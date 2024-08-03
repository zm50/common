package logcli

import (
	"common/constant"

	"k8s.io/klog/v2"
)

// Log is a struct for logging
type Log struct {}

// NewLog creates a new Log object
func NewLog(logPath string) *Log {
	return &Log{}
}

// Info logs a message at info level
func (l *Log) Info(format string, args ...any) {
	klog.InfofDepth(constant.Two, format, args...)
}

// Error logs a message at error level
func (l *Log) Error(format string, args ...any) {
	klog.ErrorfDepth(constant.Two, format, args...)
}

// Debug logs a message at debug level
func (l *Log) Debug(format string, args ...any) {
	klog.ErrorfDepth(constant.Two, format, args...)
}

// Warn logs a message at warning level
func (l *Log) Warn(format string, args ...any) {
	klog.WarningfDepth(constant.Two, format, args...)
}
