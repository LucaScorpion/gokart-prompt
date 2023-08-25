package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
)

var Node = section{
	symbol: "⬢",
	color:  ansi.Green,

	upsearchFiles: []string{
		"node_modules",
		"package.json",
		".nvmrc",
	},
	wdFiles: []string{
		"*.js",
		"*.ts",
	},

	command: []string{"node", "-v"},
	versionFn: func(output string) string {
		/*
			Example:
			v19.1.0
		*/
		return output
	},

	expectedVersionFn: func() string {
		return internal.ReadUpsearchWdFile(".nvmrc")
	},
}
