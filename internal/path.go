package internal

import (
	"os"
	"strings"
)

func Path() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	home := os.Getenv("HOME")
	if suffix, found := strings.CutPrefix(wd, home); found {
		wd = "~" + suffix
	}

	return wd
}
