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
		/*
			Example:
			go version go1.20.2 linux/amd64
		*/
		version := strings.SplitN(version, " ", 4)[2][2:]
		return ansi.Color(ansi.Cyan, " 🐹 v"+version)
	}

	return ""
}
