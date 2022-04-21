package logger

import (
	"time"
)

type ConsoleLogger struct {
	config        Config
	consoleWriter ConsoleWriter
	formatter     Formatter
	Logger
}

func (consoleLogger ConsoleLogger) IsLevelEnabled(level int) bool {
	return consoleLogger.config.IsLevelEnabled(level)
}
func (consoleLogger ConsoleLogger) IsTraceEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(TRACE)
}
func (consoleLogger ConsoleLogger) IsDebugEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(DEBUG)
}
func (consoleLogger ConsoleLogger) IsInfoEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(INFO)
}
func (consoleLogger ConsoleLogger) IsWarnEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(WARN)
}
func (consoleLogger ConsoleLogger) IsErrorEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(ERROR)
}
func (consoleLogger ConsoleLogger) IsFatalEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(FATAL)
}
func (consoleLogger ConsoleLogger) IsPanicEnabled() bool {
	return consoleLogger.config.IsLevelEnabled(PANIC)
}

func (consoleLogger ConsoleLogger) Trace(message string) {
	consoleLogger.TraceWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) TraceWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsTraceEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(TRACE), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Debug(message string) {
	consoleLogger.DebugWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) DebugWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsDebugEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(DEBUG), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Info(message string) {
	consoleLogger.InfoWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) InfoWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsInfoEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(INFO), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Warn(message string) {
	consoleLogger.WarnWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) WarnWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsWarnEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(WARN), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Error(message string) {
	consoleLogger.ErrorWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) ErrorWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsErrorEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(ERROR), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Fatal(message string) {
	consoleLogger.FatalWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) FatalWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsFatalEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(FATAL), message, metadata))
	}
}

func (consoleLogger ConsoleLogger) Panic(message string) {
	consoleLogger.PanicWithMetadata(message, nil)
}

func (consoleLogger ConsoleLogger) PanicWithMetadata(message string, metadata interface{}) {
	if consoleLogger.IsPanicEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.config.category, consoleLogger.config.levels.GetNameForValue(PANIC), message, metadata))
	}
}
