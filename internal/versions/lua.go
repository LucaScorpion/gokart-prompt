package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Lua = section{
	symbol: "🌙",
	color:  ansi.Blue,

	wdFiles: []string{
		"*.lua",
	},

	command: []string{"lua", "-v"},
	versionFn: func(output string) string {
		/*
			Example:
			Lua 5.4.6  Copyright (C) 1994-2023 Lua.org, PUC-Rio
		*/
		return "v" + strings.SplitN(output, " ", 3)[1]
	},
}
