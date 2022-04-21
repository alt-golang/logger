package logger

import (
	"fmt"
	"os"
)

type CachingConsoleWriter struct {
	Size  int
	Cache []string
	ConsoleWriter
}

func (cachingConsoleWriter CachingConsoleWriter) Append(text string) {
	cachingConsoleWriter.Cache = append(cachingConsoleWriter.Cache, text)
	if len(cachingConsoleWriter.Cache) > cachingConsoleWriter.Size {
		cachingConsoleWriter.Cache = cachingConsoleWriter.Cache[1:]
	}
}
func (cachingConsoleWriter CachingConsoleWriter) Outln(text string) {
	cachingConsoleWriter.Append(text)
	fmt.Fprintln(os.Stdout, text)
}

func (cachingConsoleWriter CachingConsoleWriter) Errln(text string) {
	cachingConsoleWriter.Append(text)
	fmt.Fprintln(os.Stderr, text)
}
