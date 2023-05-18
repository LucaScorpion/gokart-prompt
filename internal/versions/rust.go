package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var rustFiles = []string{
	"Cargo.toml",
}

func RustVersion() string {
	if _, ok := internal.UpsearchWd(rustFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("rustc", "--version"); ok {
		/*
			Example:
			rustc 1.69.0 (84c898d65 2023-04-16) (Arch Linux rust 1:1.69.0-2)
		*/
		version = strings.SplitN(version, " ", 3)[1]
		return ansi.Color(ansi.Red, " 🦀 v"+version)
	}

	return ""
}
