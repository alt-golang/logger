package logger

import (
	"fmt"
	config_pkg "github.com/alt-golang/config"
	"strings"
)

type LoggingConfig struct {
	config             config_pkg.Config
	configPath         string
	levels             Levels
	categoryLevelCache CategoryLevelCache
}

func (loggingConfig LoggingConfig) GetFormat() string {
	var format = "json"
	if loggingConfig.config.Has("logging.format") {
		result, _ := loggingConfig.config.Get("logging.format")
		format = result.(string)
	}
	return strings.ToLower(format)
}

func (loggingConfig LoggingConfig) FetchCategoryLevel(category string) {
	var path = loggingConfig.configPath
	if path == "" {
		path = DEFAULT_CONFIG_PATH
	}
	var categories = strings.Split(category, "/")
	var pathStep = path
	var root = pathStep + "./"
	if config.Has(pathStep) {
		levelString, _ := config.Get(root)
		loggingConfig.categoryLevelCache.Put(root, loggingConfig.levels.GetValueForName(strings.ToUpper(fmt.Sprint(levelString))))
	}

	for i := 0; i < len(categories); i++ {
		if i == 0 {
			pathStep = pathStep + "."
		} else {
			pathStep = pathStep + "/"
		}
		pathStep = pathStep + categories[i]
		if config.Has(pathStep) {
			levelString, _ := config.Get(pathStep)
			loggingConfig.categoryLevelCache.Put(root, loggingConfig.levels.GetValueForName(strings.ToUpper(fmt.Sprint(levelString))))
		}
	}

}
