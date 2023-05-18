package ansi

type ColorCode string

const (
	Black        ColorCode = "30"
	Red          ColorCode = "31"
	Green        ColorCode = "32"
	Yellow       ColorCode = "33"
	Blue         ColorCode = "34"
	Magenta      ColorCode = "35"
	Cyan         ColorCode = "36"
	White        ColorCode = "37"
	DefaultColor ColorCode = "39"
)

func Color(c ColorCode, s string) string {
	return mode(string(c)) + s + mode(string(DefaultColor))
}
