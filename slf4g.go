package slf4g

import (
	"fmt"
	config_pkg "github.com/alt-golang/config"
	"strings"
)

const (
	TRACE = iota
	DEBUG = iota
	INFO  = iota
	WARN  = iota
	ERROR = iota
	FATAL = iota
	PANIC = iota
)

var levelsList = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
	"PANIC",
}

var levelsMap = map[string]int{
	"TRACE": TRACE,
	"DEBUG": DEBUG,
	"INFO":  INFO,
	"WARN":  WARN,
	"ERROR": ERROR,
	"FATAL": FATAL,
	"PANIC": PANIC,
}

const DEFAULT_CONFIG_PATH = "logging.level"

var config = config_pkg.GetConfig()

var loggerCategoryCache = LoggerCategoryCache{
	cache: map[string]interface{}{},
}

func GetLoggerCategoryCache() LoggerCategoryCache {
	return loggerCategoryCache
}

var loggerFactory = LoggerFactory{
	config:              config_pkg.GetConfig(),
	loggerCategoryCache: loggerCategoryCache,
	configPath:          DEFAULT_CONFIG_PATH,
}

func GetLoggerFactory() LoggerFactory {
	return loggerFactory
}

func GetLoggerLevel(category string, configPath string, config config_pkg.Config, levels map[string]int, cache LoggerCategoryCache) int {
	var level = INFO
	var path = configPath
	if path == "" {
		path = DEFAULT_CONFIG_PATH
	}
	var categories = strings.Split(category, "/")
	var pathStep = path
	var root = pathStep + "./"
	if cache.Get(root) != nil {
		level = cache.Get(root).(int)
	} else if config.Has(pathStep) {
		levelString, _ := config.Get(root)
		cache.Put(root, levels[strings.ToUpper(fmt.Sprint(levelString))])
	}

	for i := 0; i < len(categories); i++ {
		if i == 0 {
			pathStep = pathStep + "."
		} else {
			pathStep = pathStep + "/"
		}
		pathStep = pathStep + categories[i]
		if cache.Get(pathStep) != nil {
			level = cache.Get(pathStep).(int)
		} else if config.Has(pathStep) {
			levelString, _ := config.Get(pathStep)
			cache.Put(root, levels[strings.ToUpper(fmt.Sprint(levelString))])
		}
	}
	return level
}
