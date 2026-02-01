package terminal

import (
	"gokart-prompt/internal/ansi"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Columns() int {
	cols, _ := strconv.Atoi(os.Getenv("GOKART_COLUMNS"))
	return cols
}

func RightAlign(right string) string {
	columns := Columns()
	if columns <= 0 {
		return ""
	}

	rightLen := utf8.RuneCountInString(ansi.ToPlain(right))
	return strings.Repeat(" ", Columns()-rightLen) + right
}
