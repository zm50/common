package logcli

var log ILog

// Init initializes the log client.
func Init(logPath string) {
	log = NewLog(logPath)
}

// Info logs a message with level Info.
func Info(format string, args ...any) {
	log.Info(format, args...)
}

// Error logs a message with level Error.
func Error(format string, args ...any) {
	log.Error(format, args...)
}

// Debug logs a message with level Debug.
func Debug(format string, args ...any) {
	log.Debug(format, args...)
}

// Warn logs a message with level Warn.
func Warn(format string, args ...any) {
	log.Warn(format, args...)
}
