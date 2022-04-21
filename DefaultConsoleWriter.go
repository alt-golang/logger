package logger

import (
	"fmt"
	"os"
)

type DefaultConsoleWriter struct {
	ConsoleWriter
}

func (defaultConsoleWriter DefaultConsoleWriter) Outln(text string) {
	fmt.Fprintln(os.Stdout, text)
}

func (defaultConsoleWriter DefaultConsoleWriter) Errln(text string) {
	fmt.Fprintln(os.Stderr, text)
}
