package logger

type Config struct {
	Category           string
	Levels             Levels
	CategoryLevelCache CategoryLevelCache
}

func (config Config) GetLevel() int {
	if config.CategoryLevelCache.Get(config.Category) == nil {
		return INFO
	}
	return config.CategoryLevelCache.Get(config.Category).(int)
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
