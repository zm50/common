package logcli

// ILog is the interface of log component
type ILog interface {
	Info(format string, args ...any)
	Error(format string, args ...any)
	Debug(format string, args ...any)
	Warn(format string, args ...any)
}

