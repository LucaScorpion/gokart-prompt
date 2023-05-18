package internal

import (
	"gokart-prompt/internal/ansi"
	"os"
	"path"
	"strings"
)

func Path() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	if projectDir, isProject := upsearchProjectDir(); isProject {
		dir, _ = strings.CutPrefix(dir, path.Dir(projectDir)+"/")
	} else {
		home := os.Getenv("HOME")
		if suffix, found := strings.CutPrefix(dir, home); found {
			dir = "~" + suffix
		}
	}

	return ansi.Color(ansi.Cyan, dir)
}
