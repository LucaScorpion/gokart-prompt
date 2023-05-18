package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var goFiles = []string{
	"go.mod",
}

func GoVersion() string {
	if _, ok := internal.UpsearchWd(goFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("go", "version"); ok {
		version, _, _ := strings.Cut(version[13:], " ")
		return ansi.Color(ansi.Cyan, " 🐹 v"+version)
	}

	return ""
}
