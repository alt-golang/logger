package slf4g

type LevelFunction func(string) int

type DefaultLoggerConfig struct {
	category      string
	level         int
	levelFunction LevelFunction
	levels        []string
	LoggerConfig
}

func (defaultLoggerConfig DefaultLoggerConfig) SetCategory(category string) {
	defaultLoggerConfig.category = category
}
func (defaultLoggerConfig DefaultLoggerConfig) GetLevel() int {
	if defaultLoggerConfig.levelFunction != nil {
		return defaultLoggerConfig.levelFunction(defaultLoggerConfig.category)
	}
	return defaultLoggerConfig.level
}
func (defaultLoggerConfig DefaultLoggerConfig) SetLevel(level int) {
	defaultLoggerConfig.level = level
}
func (defaultLoggerConfig DefaultLoggerConfig) IsLevelEnabled(level int) bool {
	if defaultLoggerConfig.GetLevel() <= defaultLoggerConfig.level {
		return true
	}
	return false
}
func (defaultLoggerConfig DefaultLoggerConfig) IsTraceEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(TRACE)
}
func (defaultLoggerConfig DefaultLoggerConfig) IsDebugEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(DEBUG)
}
func (defaultLoggerConfig DefaultLoggerConfig) IsInfoEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(INFO)
}
func (defaultLoggerConfig DefaultLoggerConfig) IsWarnEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(WARN)
}
func (defaultLoggerConfig DefaultLoggerConfig) IsErrorEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(ERROR)
}
func (defaultLoggerConfig DefaultLoggerConfig) IsFatalEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(FATAL)
}
func (defaultLoggerConfig DefaultLoggerConfig) IsPanicEnabled() bool {
	return defaultLoggerConfig.IsLevelEnabled(PANIC)
}
