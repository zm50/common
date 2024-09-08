package logcli

var log ILog

// Init initializes the log client.
func Init(logPath string) {
	log = NewLog(logPath)
}

// Info logs a message with level Info.
func Info(format string) {
	log.Info(format)
}

// Error logs a message with level Error.
func Error(format string) {
	log.Error(format)
}

// Warn logs a message with level Warn.
func Warn(format string) {
	log.Warn(format)
}

// Infof logs a message with level Info.
func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

// Errorf logs a message with level Error.
func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

// Warnf logs a message with level Warn.
func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}
