package slf4g

type ConsoleWriter interface {
	Outln(text string)
	Errln(text string)
}
