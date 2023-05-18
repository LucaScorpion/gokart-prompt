package ansi

const Esc = "\x1B"

func mode(mode string) string {
	return Esc + "[" + mode + "m"
}
