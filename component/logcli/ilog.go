package logcli

// ILog is the interface of log component
type ILog interface {
	Info(format string)
	Error(format string)
	Warn(format string)
	Infof(format string, args ...any)
	Errorf(format string, args ...any)
	Warnf(format string, args ...any)
}

