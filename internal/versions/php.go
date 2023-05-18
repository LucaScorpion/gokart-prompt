package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var phpFiles = []string{
	"composer.json",
}

func PhpVersion() string {
	if _, ok := internal.UpsearchWd(phpFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("php", "-v"); ok {
		parts := strings.SplitN(version, " ", 3)
		version = parts[1]

		return ansi.Color(ansi.Magenta, " 🐘 v"+version)
	}

	return ""
}
