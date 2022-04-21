package logger

import (
	"strings"
)

type DefaultFactory struct {
	loggingConfig      LoggingConfig
	format             string
	jsonFormatter      JSONFormatter
	plainTextFormatter PlainTextFormatter
	registry           Registry
	Factory
}

func (loggerFactory DefaultFactory) GetLogger(category string) Logger {

	loggingConfig.FetchCategoryLevel(category)

	config := Config{
		category: category,
		levels:   defaultLevels,
	}
	consoleLogger := ConsoleLogger{
		config:        config,
		consoleWriter: DefaultConsoleWriter{},
		formatter:     loggerFactory.GetFormatter(),
	}

	loggerFactory.registry.loggers[category] = consoleLogger
	return consoleLogger
}

func (loggerFactory DefaultFactory) GetFormatter() Formatter {
	if strings.ToLower(loggerFactory.format) == "text" {
		return loggerFactory.plainTextFormatter
	}
	return loggerFactory.jsonFormatter
}
