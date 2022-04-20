package slf4g

type Logger interface {
	GetLoggerConfig() LoggerConfig
	SetCategory(category string)
	GetLevel() int
	SetLevel(level int)
	IsLevelEnabled(level int) bool
	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool
	IsPanicEnabled() bool

	Trace(v ...any)
	Tracef(format string, v ...any)
	Traceln(v ...any)

	Debug(v ...any)
	Debugf(format string, v ...any)
	Debugln(v ...any)

	Info(v ...any)
	Infof(format string, v ...any)
	Infoln(v ...any)

	Warn(v ...any)
	Warnf(format string, v ...any)
	Warnln(v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)
	Errorln(v ...any)

	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)

	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
}
