package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"os"
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
		file, found := internal.UpsearchWd([]string{".nvmrc"})
		if !found {
			return ""
		}

		bytes, err := os.ReadFile(file)
		if err != nil {
			return ""
		}

		return string(bytes)
	},
}
