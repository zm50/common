package except

import (
	"runtime/debug"

	"github.com/zm50/common/component/logcli"
)

// Recover is a function that recovers from panic and logs the error and stack trace.
func Recover() {
	if err := recover(); err!= nil {
		logcli.Errorf("recover panic: %v", err)
		logcli.Errorf("stack trace: %s", debug.Stack())
	}
}
