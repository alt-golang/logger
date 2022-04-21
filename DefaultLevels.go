package logger

const (
	TRACE = iota
	DEBUG = iota
	INFO  = iota
	WARN  = iota
	ERROR = iota
	FATAL = iota
	PANIC = iota
)

var defaultLevelsList = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
	"PANIC",
}

var defaultLevelsMap = map[string]int{
	"TRACE": TRACE,
	"DEBUG": DEBUG,
	"INFO":  INFO,
	"WARN":  WARN,
	"ERROR": ERROR,
	"FATAL": FATAL,
	"PANIC": PANIC,
}

type DefaultLevels struct {
	asMap  map[string]int
	asList []string
}

func (defaulLevels DefaultLevels) GetValueForName(name string) int {
	return defaulLevels.asMap[name]
}
func (defaulLevels DefaultLevels) GetNameForValue(value int) string {
	return defaulLevels.asList[value]
}

var defaultLevels = DefaultLevels{
	asMap:  defaultLevelsMap,
	asList: defaultLevelsList,
}
