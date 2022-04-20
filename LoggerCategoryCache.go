package slf4g

type LoggerCategoryCache struct {
	cache map[string]interface{}
}

func (loggerCategoryCache LoggerCategoryCache) Get(category string) interface{} {
	return loggerCategoryCache.cache[category]
}

func (loggerCategoryCache LoggerCategoryCache) Put(category string, level int) {
	loggerCategoryCache.cache[category] = level
}
