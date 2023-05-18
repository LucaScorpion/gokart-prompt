package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var phpFiles = []string{
	"composer.json",
}

func PhpVersion() string {
	if _, ok := internal.UpsearchWd(phpFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("php", "-v"); ok {
		/*
			Example:
			PHP 8.2.4 (cli) (built: Mar 15 2023 15:27:52) (NTS)
			Copyright (c) The PHP Group
			Zend Engine v4.2.4, Copyright (c) Zend Technologies
		*/
		version := strings.SplitN(version, " ", 3)[1]
		return ansi.Color(ansi.Magenta, " 🐘 v"+version)
	}

	return ""
}
