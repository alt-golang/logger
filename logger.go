package logger

import (
	config_pkg "github.com/alt-golang/config"
)

type Logger interface {
	IsLevelEnabled(level int) bool
	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool
	IsPanicEnabled() bool

	Log(level int, message string)
	LogWithMetadata(level int, message string, metadata interface{})

	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	Panic(message string)

	TraceWithMetadata(message string, metadata interface{})
	DebugWithMetadata(message string, metadata interface{})
	InfoWithMetadata(message string, metadata interface{})
	WarnWithMetadata(message string, metadata interface{})
	ErrorWithMetadata(message string, metadata interface{})
	FatalWithMetadata(message string, metadata interface{})
	PanicWithMetadata(message string, metadata interface{})
}

const DEFAULT_CONFIG_PATH = "logging.level"

var config = config_pkg.GetConfig()

var categoryLevelCache = CategoryLevelCache{
	Cache: map[string]interface{}{},
}

func GetCategoryLevelCache() CategoryLevelCache {
	return categoryLevelCache
}

var loggingConfig = LoggingConfig{
	config:             config_pkg.GetConfig(),
	configPath:         DEFAULT_CONFIG_PATH,
	levels:             defaultLevels,
	categoryLevelCache: categoryLevelCache,
}

var registry = Registry{
	Loggers: map[string]Logger{},
}

var factory = DefaultFactory{
	LoggingConfig:      loggingConfig,
	Format:             loggingConfig.GetFormat(),
	PlainTextFormatter: PlainTextFormatter{},
	JSONFormatter:      JSONFormatter{},
	Registry:           registry,
}

func GetFactory() Factory {
	return factory
}

func GetRegistry() Registry {
	return registry
}

func GetLogger(category string) Logger {
	return factory.GetLogger(category)
}
