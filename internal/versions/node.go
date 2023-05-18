package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
)

var nodeFiles = []string{
	"package.json",
	"node_modules",
}

func NodeVersion() string {
	if _, ok := internal.UpsearchWd(nodeFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("node", "-v"); ok {
		return ansi.Color(ansi.Green, " ⬢ "+version)
	}

	return ""
}
