package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var Go = section{
	symbol: "🐹",
	color:  ansi.Cyan,

	upsearchFiles: []string{
		"go.mod",
	},
	wdFiles: []string{
		"*.go",
	},

	command: []string{"go", "version"},
	versionFn: func(output string) string {
		/*
			Example:
			go version go1.20.2 linux/amd64
		*/
		return "v" + strings.SplitN(output, " ", 4)[2][2:]
	},

	expectedVersionFn: func() string {
		str := internal.UpsearchWdFileContents("go.mod")
		lines := strings.Split(str, "\n")

		for _, line := range lines {
			if strings.HasPrefix(line, "go ") {
				return line[3:]
			}
		}

		return ""
	},
}
