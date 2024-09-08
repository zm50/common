package logcli

import (
	"github.com/zm50/common/constant"

	"k8s.io/klog/v2"
)

// Log is a struct for logging
type Log struct {}

// NewLog creates a new Log object
func NewLog(logPath string) *Log {
	return &Log{}
}

// Info logs a message at info level
func (l *Log) Info(format string) {
	klog.InfoDepth(constant.Two, format)
}

// Error logs a message at error level
func (l *Log) Error(format string) {
	klog.ErrorDepth(constant.Two, format)
}

// Warn logs a message at warning level
func (l *Log) Warn(format string) {
	klog.WarningDepth(constant.Two, format)
}

// Infof logs a message at info level with format
func (l *Log) Infof(format string, args ...any) {
	klog.InfofDepth(constant.Two, format, args...)
}

// Errorf logs a message at error level with format
func (l *Log) Errorf(format string, args ...any) {
	klog.ErrorfDepth(constant.Two, format, args...)
}

// Warnf logs a message at warning level with format
func (l *Log) Warnf(format string, args ...any) {
	klog.WarningfDepth(constant.Two, format, args...)
}
