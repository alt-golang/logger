package slf4g

type LoggerConfig interface {
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
}
