package versions

import (
	"gokart-powerline/internal"
	"gokart-powerline/internal/ansi"
)

var nodeFiles = []string{
	"package.json",
	"node_modules",
}

func NodeVersion() string {
	if _, ok := internal.UpsearchWd(nodeFiles); !ok {
		return ""
	}

	return ansi.Color(ansi.Green, " ⬢ v")
}
