package versions

import (
	"gokart-prompt/internal/ansi"
)

var Node = section{
	symbol: "⬢",
	color:  ansi.Green,

	upsearchFiles: []string{
		"node_modules",
		"package.json",
	},
	wdFiles: []string{
		"*.js",
		"*.ts",
	},

	command: []string{"node", "-v"},
	versionFunc: func(output string) string {
		/*
			Example:
			v19.1.0
		*/
		return output
	},
}
