package hlog

import (
	"fmt"
	"log"
	"os"
)

var (
	errorLog = log.New(os.Stdout, "\033[31;1m[ERROR] ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34;1m[INFO ] ", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(os.Stdout, "\033[33;1m[WARN ] ", log.LstdFlags|log.Lshortfile)
)

func Debug(args ...interface{}) {

}
func DebugOf(format string, args ...interface{}) {

}
func CtxDebugOf(ctx interface{}, format string, args ...interface{}) {

}

func Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	infoLog.Output(2, msg)
}
func InfoOf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	infoLog.Output(2, msg)
}

func CtxInfoOf(ctx interface{}, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	infoLog.Output(2, msg)
}

func Warn(args ...interface{}) {
	msg := fmt.Sprint(args...)
	warnLog.Output(2, msg)
}
func WarnOf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	warnLog.Output(2, msg)
}
func CtxWarnOf(ctx interface{}, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	warnLog.Output(2, msg)
}

func Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	errorLog.Output(2, msg)
}
func ErrorOf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	errorLog.Output(2, msg)
}
func CtxErrorOf(ctx interface{}, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	errorLog.Output(2, msg)
}

func Fatal(args ...interface{}) {

}
func FatalOf(format string, args ...interface{}) {

}
func CtxFatalOf(ctx interface{}, format string, args ...interface{}) {

}

type logImpl struct {
}

func (l *logImpl) Debug(args ...interface{}) {
	Debug(args...)
}
func (l *logImpl) DebugOf(format string, args ...interface{}) {
	DebugOf(format, args...)
}
func (l *logImpl) CtxDebugOf(ctx interface{}, format string, args ...interface{}) {
	CtxDebugOf(ctx, format, args...)
}

func (l *logImpl) Info(args ...interface{}) {
	Info(args...)
}
func (l *logImpl) InfoOf(format string, args ...interface{}) {
	InfoOf(format, args...)
}
func (l *logImpl) CtxInfoOf(ctx interface{}, format string, args ...interface{}) {
	CtxInfoOf(ctx, format, args...)
}

func (l *logImpl) Warn(args ...interface{}) {
	Warn(args...)
}
func (l *logImpl) WarnOf(format string, args ...interface{}) {
	WarnOf(format, args...)
}
func (l *logImpl) CtxWarnOf(ctx interface{}, format string, args ...interface{}) {
	CtxWarnOf(ctx, format, args...)
}

func (l *logImpl) Error(args ...interface{}) {
	Error(args...)
}
func (l *logImpl) ErrorOf(format string, args ...interface{}) {
	ErrorOf(format, args...)
}
func (l *logImpl) CtxErrorOf(ctx interface{}, format string, args ...interface{}) {
	CtxErrorOf(ctx, format, args...)
}

func (l *logImpl) Fatal(args ...interface{}) {
	Fatal(args...)
}
func (l *logImpl) FatalOf(format string, args ...interface{}) {
	FatalOf(format, args...)
}
func (l *logImpl) CtxFatalOf(ctx interface{}, format string, args ...interface{}) {
	CtxFatalOf(ctx, format, args...)
}
