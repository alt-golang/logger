package logger

import (
	"time"
)

type ConsoleLogger struct {
	Config        Config
	ConsoleWriter ConsoleWriter
	Formatter     Formatter
	Logger
}

func (consoleLogger ConsoleLogger) IsLevelEnabled(level int) bool {
	return consoleLogger.Config.IsLevelEnabled(level)
}
func (consoleLogger ConsoleLogger) IsTraceEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(TRACE)
}
func (consoleLogger ConsoleLogger) IsDebugEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(DEBUG)
}
func (consoleLogger ConsoleLogger) IsInfoEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(INFO)
}
func (consoleLogger ConsoleLogger) IsWarnEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(WARN)
}
func (consoleLogger ConsoleLogger) IsErrorEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(ERROR)
}
func (consoleLogger ConsoleLogger) IsFatalEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(FATAL)
}
func (consoleLogger ConsoleLogger) IsPanicEnabled() bool {
	return consoleLogger.Config.IsLevelEnabled(PANIC)
}

func (consoleLogger ConsoleLogger) Trace(message string) {
	consoleLogger.TraceWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) TraceWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsTraceEnabled() {
		consoleLogger.ConsoleWriter.Outln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(TRACE), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Debug(message string) {
	consoleLogger.DebugWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) DebugWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsDebugEnabled() {
		consoleLogger.ConsoleWriter.Outln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(DEBUG), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Info(message string) {
	consoleLogger.InfoWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) InfoWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsInfoEnabled() {
		consoleLogger.ConsoleWriter.Outln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(INFO), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Warn(message string) {
	consoleLogger.WarnWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) WarnWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsWarnEnabled() {
		consoleLogger.ConsoleWriter.Errln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(WARN), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Error(message string) {
	consoleLogger.ErrorWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) ErrorWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsErrorEnabled() {
		consoleLogger.ConsoleWriter.Errln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(ERROR), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Fatal(message string) {
	consoleLogger.FatalWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) FatalWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsFatalEnabled() {
		consoleLogger.ConsoleWriter.Errln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(FATAL), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Panic(message string) {
	consoleLogger.PanicWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) PanicWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsPanicEnabled() {
		consoleLogger.ConsoleWriter.Errln(consoleLogger.Formatter.Format(time.Now(), consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(PANIC), message, metadata))
	}
}
