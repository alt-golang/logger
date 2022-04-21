package logger

type CategoryLevelCache struct {
	Cache map[string]interface{}
}

func (categoryLevelCache CategoryLevelCache) Get(category string) interface{} {
	return categoryLevelCache.Cache[category]
}

func (categoryLevelCache CategoryLevelCache) Put(category string, level int) {
	categoryLevelCache.Cache[category] = level
}
