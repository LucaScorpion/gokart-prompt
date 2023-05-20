package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Php = section{
	symbol: "🐘",
	color:  ansi.Magenta,

	upsearchFiles: []string{
		"composer.json",
	},

	command: []string{"php", "-v"},
	versionFunc: func(output string) string {
		/*
			Example:
			PHP 8.2.4 (cli) (built: Mar 15 2023 15:27:52) (NTS)
			Copyright (c) The PHP Group
			Zend Engine v4.2.4, Copyright (c) Zend Technologies
		*/
		return "v" + strings.SplitN(output, " ", 3)[1]
	},
}
