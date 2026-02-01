package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"gokart-prompt/internal/semVer"
	"path"
)

type section struct {
	symbol            string
	color             ansi.ColorCode
	upsearchFiles     []string
	wdFiles           []string
	command           []string
	versionFn         func(output string) string
	expectedVersionFn func() string
}

func (s section) version(wdFiles []string) string {
	if !s.display(wdFiles) {
		return ""
	}

	if output, ok := internal.Command(s.command[0], s.command[1:]...); ok {
		curVersion := s.versionFn(output)
		str := s.symbol + " " + curVersion

		// Check if we know the expected version, and if that matches the found version.
		if s.expectedVersionFn != nil && !semVer.Matches(curVersion, s.expectedVersionFn()) {
			str += " ⚠️ "
		}

		return ansi.Color(s.color, str)
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
