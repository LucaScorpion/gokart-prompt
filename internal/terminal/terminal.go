package terminal

import (
	"os"
	"strconv"
)

func Columns() int {
	cols, _ := strconv.Atoi(os.Getenv("GOKART_COLUMNS"))
	return cols
}

func IsZsh() bool {
	return os.Getenv("GOKART_SHELL") == "zsh"
}
