package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Java = section{
	symbol: "☕",
	color:  ansi.Cyan,

	upsearchFiles: []string{
		"pom.xml",
		"build.gradle",
		"settings.gradle",
		"build.xml",
	},

	command: []string{"java", "--version"},
	versionFunc: func(output string) string {
		/*
			Example:
			openjdk 19.0.2 2023-01-17
			OpenJDK Runtime Environment (build 19.0.2+7)
			OpenJDK 64-Bit Server VM (build 19.0.2+7, mixed mode)
		*/
		return "v" + strings.SplitN(output, " ", 3)[1]
	},
}
