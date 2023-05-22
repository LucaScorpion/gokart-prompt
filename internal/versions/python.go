package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Python = section{
	symbol: "🐍",
	color:  ansi.Yellow,

	upsearchFiles: []string{
		"Pipfile",
		"pyproject.toml",
		"requirements.txt",
	},
	wdFiles: []string{
		"*.py",
	},

	command: []string{"python", "--version"},
	versionFunc: func(output string) string {
		/*
			Example:
			Python 3.10.10
		*/
		return "v" + strings.Split(output, " ")[1]
	},
}
