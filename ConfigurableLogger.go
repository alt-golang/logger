package slf4g

type ConfigurableLogger struct {
	logger Logger
	Logger
}

func (configurableLogger ConfigurableLogger) GetLoggerConfig() LoggerConfig {
	return configurableLogger.logger.GetLoggerConfig()
}
func (configurableLogger ConfigurableLogger) SetCategory(category string) {
	configurableLogger.logger.SetCategory(category)
}
func (configurableLogger ConfigurableLogger) GetLevel() int {
	return configurableLogger.logger.GetLevel()
}
func (configurableLogger ConfigurableLogger) SetLevel(level int) {
	configurableLogger.logger.SetLevel(level)
}
func (configurableLogger ConfigurableLogger) IsLevelEnabled(level int) bool {
	return configurableLogger.logger.IsLevelEnabled(level)
}
func (configurableLogger ConfigurableLogger) IsTraceEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(TRACE)
}
func (configurableLogger ConfigurableLogger) IsDebugEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(DEBUG)
}
func (configurableLogger ConfigurableLogger) IsInfoEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(INFO)
}
func (configurableLogger ConfigurableLogger) IsWarnEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(WARN)
}
func (configurableLogger ConfigurableLogger) IsErrorEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(ERROR)
}
func (configurableLogger ConfigurableLogger) IsFatalEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(FATAL)
}
func (configurableLogger ConfigurableLogger) IsPanicEnabled() bool {
	return configurableLogger.logger.IsLevelEnabled(PANIC)
}

func (configurableLogger ConfigurableLogger) Trace(v ...any) {
	configurableLogger.logger.Trace(v)
}
func (configurableLogger ConfigurableLogger) Tracef(format string, v ...any) {
	configurableLogger.logger.Tracef(format, v)
}
func (configurableLogger ConfigurableLogger) Traceln(v ...any) {
	configurableLogger.logger.Traceln(v)
}

func (configurableLogger ConfigurableLogger) Debug(v ...any) {
	configurableLogger.logger.Debug(v)
}
func (configurableLogger ConfigurableLogger) Debugf(format string, v ...any) {
	configurableLogger.logger.Debugf(format, v)
}
func (configurableLogger ConfigurableLogger) Debugln(v ...any) {
	configurableLogger.logger.Debugln(v)
}

func (configurableLogger ConfigurableLogger) Info(v ...any) {
	configurableLogger.logger.Info(v)
}
func (configurableLogger ConfigurableLogger) Infof(format string, v ...any) {
	configurableLogger.logger.Infof(format, v)
}
func (configurableLogger ConfigurableLogger) Infoln(v ...any) {
	configurableLogger.logger.Infoln(v)
}

func (configurableLogger ConfigurableLogger) Warn(v ...any) {
	configurableLogger.logger.Warn(v)
}
func (configurableLogger ConfigurableLogger) Warnf(format string, v ...any) {
	configurableLogger.logger.Warnf(format, v)
}
func (configurableLogger ConfigurableLogger) Warnln(v ...any) {
	configurableLogger.logger.Warnln(v)
}

func (configurableLogger ConfigurableLogger) Error(v ...any) {
	configurableLogger.logger.Error(v)
}
func (configurableLogger ConfigurableLogger) Errorf(format string, v ...any) {
	configurableLogger.logger.Errorf(format, v)
}
func (configurableLogger ConfigurableLogger) Errorln(v ...any) {
	configurableLogger.logger.Errorln(v)
}

func (configurableLogger ConfigurableLogger) Fatal(v ...any) {
	configurableLogger.logger.Fatal(v)
}
func (configurableLogger ConfigurableLogger) Fatalf(format string, v ...any) {
	configurableLogger.logger.Fatalf(format, v)
}
func (configurableLogger ConfigurableLogger) Fatalln(v ...any) {
	configurableLogger.logger.Fatalln(v)
}

func (configurableLogger ConfigurableLogger) Panic(v ...any) {
	configurableLogger.logger.Panic(v)
}
func (configurableLogger ConfigurableLogger) Panicf(format string, v ...any) {
	configurableLogger.logger.Panicf(format, v)
}
func (configurableLogger ConfigurableLogger) Panicln(v ...any) {
	configurableLogger.logger.Panicln(v)
}
