package internal

import (
	"gokart-prompt/internal/ansi"
	"os"
	"path"
	"strings"
)

func Path() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	if isProjectDir(wd) {
		wd = path.Base(wd)
	} else {
		home := os.Getenv("HOME")
		if suffix, found := strings.CutPrefix(wd, home); found {
			wd = "~" + suffix
		}
	}

	return ansi.Color(ansi.Cyan, wd)
}
