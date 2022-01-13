package hlog

type Logger interface {
	Debug(args ...interface{})
	DebugOf(format string, args ...interface{})
	CtxDebugOf(ctx interface{}, format string, args ...interface{})

	Info(args ...interface{})
	InfoOf(format string, args ...interface{})
	CtxInfoOf(ctx interface{}, format string, args ...interface{})

	Warn(args ...interface{})
	WarnOf(format string, args ...interface{})
	CtxWarnOf(ctx interface{}, format string, args ...interface{})

	Error(args ...interface{})
	ErrorOf(format string, args ...interface{})
	CtxErrorOf(ctx interface{}, format string, args ...interface{})

	Fatal(args ...interface{})
	FatalOf(format string, args ...interface{})
	CtxFatalOf(ctx interface{}, format string, args ...interface{})
}
