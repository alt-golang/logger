package logger

import "time"

type Formatter interface {
	Format(t time.Time, category string, level string, message string, meta interface{}) string
}
