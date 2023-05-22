package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"path"
)

type section struct {
	symbol        string
	color         ansi.ColorCode
	upsearchFiles []string
	wdFiles       []string
	command       []string
	versionFunc   func(output string) string
}

func (s section) version(wdFiles []string) string {
	if !s.display(wdFiles) {
		return ""
	}

	if output, ok := internal.Command(s.command[0], s.command[1:]...); ok {
		return ansi.Color(s.color, " "+s.symbol+" "+s.versionFunc(output))
	}

	return ""
}

func (s section) display(wdFiles []string) bool {
	for _, file := range wdFiles {
		for _, pattern := range s.wdFiles {
			if ok, _ := path.Match(pattern, file); ok {
				return true
			}
		}
	}

	if _, ok := internal.UpsearchWd(s.upsearchFiles); ok {
		return true
	}

	return false
}
