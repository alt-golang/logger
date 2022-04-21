package logger

type Config struct {
	category           string
	levels             Levels
	categoryLevelCache CategoryLevelCache
}

func (config Config) GetLevel() int {
	if config.categoryLevelCache.Get(config.category) == nil {
		return INFO
	}
	return config.categoryLevelCache.Get(config.category).(int)
}
func (config Config) IsLevelEnabled(level int) bool {
	if config.GetLevel() <= level {
		return true
	}
	return false
}
func (config Config) IsTraceEnabled() bool {
	return config.IsLevelEnabled(TRACE)
}
func (config Config) IsDebugEnabled() bool {
	return config.IsLevelEnabled(DEBUG)
}
func (config Config) IsInfoEnabled() bool {
	return config.IsLevelEnabled(INFO)
}
func (config Config) IsWarnEnabled() bool {
	return config.IsLevelEnabled(WARN)
}
func (config Config) IsErrorEnabled() bool {
	return config.IsLevelEnabled(ERROR)
}
func (config Config) IsFatalEnabled() bool {
	return config.IsLevelEnabled(FATAL)
}
func (config Config) IsPanicEnabled() bool {
	return config.IsLevelEnabled(PANIC)
}
