package loggerv1

import (
	"fmt"

	configsv2 "github.com/limboware/pkg/configs/v2"
	errnov1 "github.com/rejchev/errno"
	"github.com/untrustedmodders/go-plugify"
)

type Logger struct {
	flags LoggerFlag
}

var l Logger

func Get() *Logger {
	return &l
}

func (x *Logger) Init() errnov1.Code {
	x.flags = None

	if !Parse(configsv2.Get().ValueA("logs", "sys,errs,warns"), &x.flags) {
		return errnov1.EINVAL
	}

	return errnov1.OK
}

func (x *Logger) severity(v LoggerFlag) plugify.Severity {
	switch v {
	case PrintLogs:
		return plugify.Info

	case PrintDebug:
		return plugify.Debug

	case PrintErrs:
		return plugify.Error

	case PrintWarns:
		return plugify.Warning

	case PrintSystem:
		return plugify.Info

	default:
		return plugify.Unknown
	}
}

func (x *Logger) Log(buff string, v ...any) {
	if (x.flags & PrintLogs) != 0 {
		plugify.Log(fmt.Sprintf(buff, v...), x.severity(PrintLogs), "", 0)
	}
}

func (x *Logger) Debug(buff string, v ...any) {
	if (x.flags & PrintDebug) != 0 {
		plugify.Log(fmt.Sprintf(buff, v...), x.severity(PrintDebug), "", 0)
	}
}

func (x *Logger) System(buff string, v ...any) {
	if (x.flags & PrintSystem) != 0 {
		plugify.Log(fmt.Sprintf(buff, v...), x.severity(PrintSystem), "", 0)
	}
}

func (x *Logger) Warn(buff string, v ...any) {
	if (x.flags & PrintWarns) != 0 {
		plugify.Log(fmt.Sprintf(buff, v...), x.severity(PrintWarns), "", 0)
	}
}

func (x *Logger) Err(buff string, v ...any) {
	if (x.flags & PrintErrs) != 0 {
		plugify.Log(fmt.Sprintf(buff, v...), x.severity(PrintErrs), "", 0)
	}
}
