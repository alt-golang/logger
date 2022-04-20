package slf4g

import (
	config_pkg "github.com/alt-golang/config"
	"strings"
)

type LoggerFactory struct {
	config              config_pkg.Config
	loggerCategoryCache LoggerCategoryCache
	configPath          string
}

func (loggerFactory LoggerFactory) GetLogger(category string) Logger {

	defaultLoggerConfig := DefaultLoggerConfig{
		category:      category,
		level:         loggerFactory.GetLevel(category),
		levelFunction: loggerFactory.GetLevel,
		levels:        levelsList,
	}
	consoleLogger := ConsoleLogger{
		defaultLoggerConfig: defaultLoggerConfig,
		consoleWriter:       DefaultConsoleWriter{},
		formatter:           loggerFactory.GetFormatter(),
	}
	configurableLogger := ConfigurableLogger{
		logger: consoleLogger,
	}
	return configurableLogger
}

func (loggerFactory LoggerFactory) GetLevel(category string) int {
	configPath := loggerFactory.configPath
	if configPath == "" {
		configPath = DEFAULT_CONFIG_PATH
	}
	return GetLoggerLevel(category, configPath, loggerFactory.config, levelsMap, loggerFactory.loggerCategoryCache)
}

func (loggerFactory LoggerFactory) GetFormatter() Formatter {
	var format = "json"
	if loggerFactory.config.Has("logging.format") {
		result, _ := loggerFactory.config.Get("logging.format")
		format = result.(string)
	}
	if strings.ToLower(format) == "text" {
		return PlainTextFormatter{}
	}
	return PlainTextFormatter{}
}
