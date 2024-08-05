package except

import (
	"runtime/debug"

	"github.com/zm50/common/component/logcli"
)

func Recover() {
	if err := recover(); err!= nil {
		logcli.Error("recover panic: %v", err)
		logcli.Error("stack trace: %s", debug.Stack())
	}
}
