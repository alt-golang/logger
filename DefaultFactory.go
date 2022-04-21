package logger

import (
	"strings"
)

type DefaultFactory struct {
	LoggingConfig      LoggingConfig
	Format             string
	JSONFormatter      JSONFormatter
	PlainTextFormatter PlainTextFormatter
	Registry           Registry
	Factory
}

func (loggerFactory DefaultFactory) GetLogger(category string) Logger {

	loggingConfig.FetchCategoryLevel(category)

	config := Config{
		Category: category,
		Levels:   defaultLevels,
	}
	consoleLogger := ConsoleLogger{
		Config:        config,
		ConsoleWriter: DefaultConsoleWriter{},
		Formatter:     loggerFactory.GetFormatter(),
	}

	loggerFactory.Registry.Loggers[category] = consoleLogger
	return consoleLogger
}

func (loggerFactory DefaultFactory) GetFormatter() Formatter {
	if strings.ToLower(loggerFactory.Format) == "text" {
		return loggerFactory.PlainTextFormatter
	}
	return loggerFactory.JSONFormatter
}
