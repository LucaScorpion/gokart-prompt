package internal

import (
	"os"
	"strconv"
)

func Columns() int {
	cols, _ := strconv.Atoi(os.Getenv("COLUMNS"))
	return cols
}
