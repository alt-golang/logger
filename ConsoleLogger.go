package slf4g

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	defaultLoggerConfig DefaultLoggerConfig
	consoleWriter       ConsoleWriter
	formatter           Formatter
	Logger
}

func (consoleLogger ConsoleLogger) GetLoggerConfig() LoggerConfig {
	return consoleLogger.defaultLoggerConfig
}
func (consoleLogger ConsoleLogger) SetCategory(category string) {
	consoleLogger.defaultLoggerConfig.SetCategory(category)
}
func (consoleLogger ConsoleLogger) GetLevel() int {
	return consoleLogger.defaultLoggerConfig.GetLevel()
}
func (consoleLogger ConsoleLogger) SetLevel(level int) {
	consoleLogger.defaultLoggerConfig.SetLevel(level)
}
func (consoleLogger ConsoleLogger) IsLevelEnabled(level int) bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(level)
}
func (consoleLogger ConsoleLogger) IsTraceEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(TRACE)
}
func (consoleLogger ConsoleLogger) IsDebugEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(DEBUG)
}
func (consoleLogger ConsoleLogger) IsInfoEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(INFO)
}
func (consoleLogger ConsoleLogger) IsWarnEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(WARN)
}
func (consoleLogger ConsoleLogger) IsErrorEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(ERROR)
}
func (consoleLogger ConsoleLogger) IsFatalEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(FATAL)
}
func (consoleLogger ConsoleLogger) IsPanicEnabled() bool {
	return consoleLogger.defaultLoggerConfig.IsLevelEnabled(PANIC)
}

func (consoleLogger ConsoleLogger) Trace(v ...any) {
	if consoleLogger.IsTraceEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[TRACE], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Tracef(format string, v ...any) {
	if consoleLogger.IsTraceEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[TRACE], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Traceln(v ...any) {
	if consoleLogger.IsTraceEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[TRACE], fmt.Sprintln(v...), nil))
	}
}

func (consoleLogger ConsoleLogger) Debug(v ...any) {
	if consoleLogger.IsDebugEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[DEBUG], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Debugf(format string, v ...any) {
	if consoleLogger.IsDebugEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[DEBUG], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Debugln(v ...any) {
	if consoleLogger.IsDebugEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[DEBUG], fmt.Sprintln(v...), nil))
	}
}

func (consoleLogger ConsoleLogger) Info(v ...any) {
	if consoleLogger.IsInfoEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[INFO], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Infof(format string, v ...any) {
	if consoleLogger.IsInfoEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[INFO], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Infoln(v ...any) {
	if consoleLogger.IsInfoEnabled() {
		consoleLogger.consoleWriter.Outln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[INFO], fmt.Sprintln(v...), nil))
	}
}

func (consoleLogger ConsoleLogger) Warn(v ...any) {
	if consoleLogger.IsWarnEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[WARN], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Warnf(format string, v ...any) {
	if consoleLogger.IsWarnEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[WARN], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Warnln(v ...any) {
	if consoleLogger.IsWarnEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[WARN], fmt.Sprintln(v...), nil))
	}
}

func (consoleLogger ConsoleLogger) Error(v ...any) {
	if consoleLogger.IsErrorEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[ERROR], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Errorf(format string, v ...any) {
	if consoleLogger.IsErrorEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[ERROR], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Errorln(v ...any) {
	if consoleLogger.IsErrorEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[ERROR], fmt.Sprintln(v...), nil))
	}
}

func (consoleLogger ConsoleLogger) Fatal(v ...any) {
	if consoleLogger.IsFatalEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[FATAL], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Fatalf(format string, v ...any) {
	if consoleLogger.IsFatalEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[FATAL], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Fatalln(v ...any) {
	if consoleLogger.IsFatalEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[FATAL], fmt.Sprintln(v...), nil))
	}
}

func (consoleLogger ConsoleLogger) Panic(v ...any) {
	if consoleLogger.IsPanicEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[PANIC], fmt.Sprintln(v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Panicf(format string, v ...any) {
	if consoleLogger.IsPanicEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[PANIC], fmt.Sprintf(format, v...), nil))
	}
}
func (consoleLogger ConsoleLogger) Panicln(v ...any) {
	if consoleLogger.IsPanicEnabled() {
		consoleLogger.consoleWriter.Errln(consoleLogger.formatter.Format(time.Now(), consoleLogger.defaultLoggerConfig.category, consoleLogger.defaultLoggerConfig.levels[PANIC], fmt.Sprintln(v...), nil))
	}
}
