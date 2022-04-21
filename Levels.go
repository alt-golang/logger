package logger

type Levels interface {
	GetValueForName(name string) int
	GetNameForValue(value int) string
}
