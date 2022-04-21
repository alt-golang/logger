package logger

import (
	"fmt"
	"time"
)

type PlainTextFormatter struct {
	Formatter
}

func (plainTextFormatter PlainTextFormatter) Format(t time.Time, category string, level string, message string, meta interface{}) string {
	result := t.Format("2006-Jan-02 Monday 03:04:05.999 PM MST -07:00") + ":" + category + ":" + level + ":" + message

	if meta != nil {
		result = result + fmt.Sprint(meta)
	}
	return result
}
