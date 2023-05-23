package ansi

import "os"

const esc = "\x1B"
const reset = "0"

// See: https://stackoverflow.com/a/19731390/2831123
const zshZeroWidthStart = "%{"
const zshZeroWidthEnd = "%}"

// See: https://unix.stackexchange.com/a/504152/176402
const bashZeroWidthStart = "\x01"
const bashZeroWidthEnd = "\x02"

func mode(mode string) string {
	return zeroWidth(esc + "[" + mode + "m")
}

func Reset() string {
	return mode(reset)
}

func zeroWidth(s string) string {
	if os.Getenv("GOKART_SHELL") == "zsh" {
		return zshZeroWidthStart + s + zshZeroWidthEnd
	}
	return bashZeroWidthStart + s + bashZeroWidthEnd
}
