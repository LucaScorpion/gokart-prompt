package ansi

import (
	"gokart-prompt/internal/terminal"
	"regexp"
)

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
	return zeroWidthStart() + s + zeroWidthEnd()
}

func zeroWidthStart() string {
	if terminal.IsZsh() {
		return zshZeroWidthStart
	}
	return bashZeroWidthStart
}

func zeroWidthEnd() string {
	if terminal.IsZsh() {
		return zshZeroWidthEnd
	}
	return bashZeroWidthEnd
}

func ToPlain(s string) string {
	return regexp.
		MustCompile(regexp.QuoteMeta(zeroWidthStart())+".*?"+regexp.QuoteMeta(zeroWidthEnd())).
		ReplaceAllString(s, "")
}
