package versions

import (
	"gokart-prompt/internal/ansi"
)

var Dotnet = section{
	symbol: "\uE77F", // 
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
	versionFunc: func(output string) string {
		/*
			Example:
			6.0.114
		*/
		return "v" + output
	},
}
