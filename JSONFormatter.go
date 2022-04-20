package slf4g

import (
	"fmt"
	"time"
)

type JSONFormatter struct {
	Formatter
}

func (jsonFormatter JSONFormatter) Format(t time.Time, category string, level string, message string, meta interface{}) string {
	result := fmt.Sprint(t) + ":" + category + ":" + level + ":" + message

	if meta != nil {
		result = result + fmt.Sprint(meta)
	}
	return result
}
