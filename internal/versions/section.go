package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
)

type Section struct {
	symbol        string
	color         ansi.ColorCode
	upsearchFiles []string
	command       []string
	versionFunc   func(output string) string
}

func (s Section) GetVersion() string {
	if _, ok := internal.UpsearchWd(s.upsearchFiles); !ok {
		return ""
	}

	if output, ok := internal.Command(s.command[0], s.command[1:]...); ok {
		return ansi.Color(s.color, " "+s.symbol+" "+s.versionFunc(output))
	}

	return ""
}
