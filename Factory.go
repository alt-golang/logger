package logger

type Factory interface {
	GetLogger(category string) Logger
}
