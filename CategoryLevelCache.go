package logger

type CategoryLevelCache struct {
	cache map[string]interface{}
}

func (categoryLevelCache CategoryLevelCache) Get(category string) interface{} {
	return categoryLevelCache.cache[category]
}

func (categoryLevelCache CategoryLevelCache) Put(category string, level int) {
	categoryLevelCache.cache[category] = level
}
