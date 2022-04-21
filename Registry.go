package logger

type Registry struct {
	loggers map[string]Logger
}
