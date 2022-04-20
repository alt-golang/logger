package slf4g

import (
	"fmt"
	"time"
)

type PlainTextFormatter struct {
	Formatter
}

func (plainTextFormatter PlainTextFormatter) Format(t time.Time, category string, level string, message string, meta interface{}) string {
	result := fmt.Sprint(t) + ":" + category + ":" + level + ":" + message

	if meta != nil {
		result = result + fmt.Sprint(meta)
	}
	return result
}
