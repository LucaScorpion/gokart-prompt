package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Rust = section{
	symbol: "🦀",
	color:  ansi.Red,

	upsearchFiles: []string{
		"Cargo.toml",
	},
	wdFiles: []string{
		"*.rs",
	},

	command: []string{"rustc", "--version"},
	versionFn: func(output string) string {
		/*
			Example:
			rustc 1.69.0 (84c898d65 2023-04-16) (Arch Linux rust 1:1.69.0-2)
		*/
		return "v" + strings.SplitN(output, " ", 3)[1]
	},
}
