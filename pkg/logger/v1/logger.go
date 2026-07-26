package loggerv1

import (
	"log"

	configsv2 "github.com/limboware/pkg/configs/v2"
	errnov1 "github.com/rejchev/errno"
)

type Logger struct {
	base  *log.Logger
	flags LoggerFlag
}

var l Logger

func Get() *Logger {
	return &l
}

func (x *Logger) Init() errnov1.Code {
	x.base = log.Default()
	x.flags = None

	if !Parse(configsv2.Get().ValueA("logs", "sys,errs,warns"), &x.flags) {
		return errnov1.EINVAL
	}

	return errnov1.OK
}

func (x *Logger) Closer() *log.Logger {
	return x.base
}

func (x *Logger) Log(buff string, v ...any) {
	if x.base != nil && (x.flags&PrintLogs) != 0 {
		x.Closer().Printf(buff+"\n", v...)
	}
}

func (x *Logger) Debug(buff string, v ...any) {
	if x.base != nil && (x.flags&PrintDebug) != 0 {
		x.Closer().Printf(buff+"\n", v...)
	}
}

func (x *Logger) System(buff string, v ...any) {
	if x.base != nil && (x.flags&PrintSystem) != 0 {
		x.Closer().Printf(buff+"\n", v...)
	}
}

func (x *Logger) Warn(buff string, v ...any) {
	if x.base != nil && (x.flags&PrintWarns) != 0 {
		x.Closer().Printf(buff+"\n", v...)
	}
}

func (x *Logger) Err(buff string, v ...any) {
	if x.base != nil && (x.flags&PrintErrs) != 0 {
		x.Closer().Printf(buff+"\n", v...)
	}
}
