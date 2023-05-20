package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Ruby = section{
	symbol: "💎",
	color:  ansi.Red,

	upsearchFiles: []string{
		"Gemfile",
		"Rakefile",
	},

	command: []string{"ruby", "-v"},
	versionFunc: func(output string) string {
		/*
			Example:
			ruby 3.0.5p211 (2022-11-24 revision ba5cf0f7c5) [x86_64-linux]
		*/
		return "v" + strings.SplitN(output, " ", 3)[1]
	},
}
