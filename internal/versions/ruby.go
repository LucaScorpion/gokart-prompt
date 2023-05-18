package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var rubyFiles = []string{
	"Gemfile",
	"Rakefile",
}

func RubyVersion() string {
	if _, ok := internal.UpsearchWd(rubyFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("ruby", "-v"); ok {
		/*
			Example:
			ruby 3.0.5p211 (2022-11-24 revision ba5cf0f7c5) [x86_64-linux]
		*/
		version = strings.SplitN(version, " ", 3)[1]
		return ansi.Color(ansi.Red, " 💎 v"+version)
	}

	return ""
}
