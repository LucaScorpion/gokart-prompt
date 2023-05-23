package ansi

const esc = "\x1B"
const reset = "0"

// See: https://stackoverflow.com/a/19731390/2831123
const zshZeroWidthStart = "%{"
const zshZeroWidthEnd = "%}"

func mode(mode string) string {
	return zeroWidth(esc + "[" + mode + "m")
}

func Reset() string {
	return mode(reset)
}

func zeroWidth(s string) string {
	return zshZeroWidthStart + s + zshZeroWidthEnd
}
