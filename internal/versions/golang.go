package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Go = section{
	symbol: "🐹",
	color:  ansi.Cyan,

	upsearchFiles: []string{
		"go.mod",
	},

	command: []string{"go", "version"},
	versionFunc: func(output string) string {
		/*
			Example:
			go version go1.20.2 linux/amd64
		*/
		return "v" + strings.SplitN(output, " ", 4)[2][2:]
	},
}
