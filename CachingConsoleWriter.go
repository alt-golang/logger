package slf4g

import (
	"fmt"
	"os"
)

type CachingConsoleWriter struct {
	size  int
	cache []string
	ConsoleWriter
}

func (cachingConsoleWriter CachingConsoleWriter) Cache(text string) {
	cachingConsoleWriter.cache = append(cachingConsoleWriter.cache, text)
	if len(cachingConsoleWriter.cache) > cachingConsoleWriter.size {
		cachingConsoleWriter.cache = cachingConsoleWriter.cache[1:]
	}
}
func (cachingConsoleWriter CachingConsoleWriter) Outln(text string) {
	cachingConsoleWriter.Cache(text)
	fmt.Fprintln(os.Stdout, text)
}

func (cachingConsoleWriter CachingConsoleWriter) Errln(text string) {
	cachingConsoleWriter.Cache(text)
	fmt.Fprintln(os.Stderr, text)
}
