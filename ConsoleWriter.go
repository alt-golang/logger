package logger

type ConsoleWriter interface {
	Outln(text string)
	Errln(text string)
}
