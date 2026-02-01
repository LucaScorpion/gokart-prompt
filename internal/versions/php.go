package versions

import (
	"encoding/json"
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var Php = section{
	symbol: "🐘",
	color:  ansi.Magenta,

	upsearchFiles: []string{
		"composer.json",
	},
	wdFiles: []string{
		"*.php",
	},

	command: []string{"php", "-v"},
	versionFn: func(output string) string {
		/*
			Example:
			PHP 8.2.4 (cli) (built: Mar 15 2023 15:27:52) (NTS)
			Copyright (c) The PHP Group
			Zend Engine v4.2.4, Copyright (c) Zend Technologies
		*/
		return "v" + strings.SplitN(output, " ", 3)[1]
	},

	expectedVersionFn: func() string {
		str := internal.UpsearchWdFileContents("composer.json")

		var data composerJson
		if err := json.Unmarshal([]byte(str), &data); err != nil {
			return ""
		}

		return data.Require.Php
	},
}

type composerJson struct {
	Require composerJsonRequire `json:"require"`
}

type composerJsonRequire struct {
	Php string `json:"php"`
}
