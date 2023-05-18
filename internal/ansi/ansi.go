package ansi

const esc = "\x1B"

// See: https://stackoverflow.com/a/19731390/2831123
const zeroWidthStart = "%{"
const zeroWidthEnd = "%}"

func mode(mode string) string {
	return zeroWidthStart + esc + "[" + mode + "m" + zeroWidthEnd
}
