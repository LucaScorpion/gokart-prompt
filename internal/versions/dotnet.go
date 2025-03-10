package versions

import (
	"gokart-prompt/internal/ansi"
)

var Dotnet = section{
	symbol: "🟪", // Purple square
	color:  ansi.Magenta,

	upsearchFiles: []string{
		"global.json",
	},
	wdFiles: []string{
		"*.fs",
		"*.fsproj",
		"*.cs",
		"*.csproj",
		"*.sln",
	},

	command: []string{"dotnet", "--version"},
	versionFn: func(output string) string {
		/*
			Example:
			6.0.114
		*/
		return "v" + output
	},
}
