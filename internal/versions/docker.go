package versions

import (
	"gokart-prompt/internal/ansi"
	"strings"
)

var Docker = section{
	symbol: "🐳",
	color:  ansi.Blue,

	upsearchFiles: []string{
		"Dockerfile",
		"docker-compose.yml",
		"docker-compose.yaml",
	},

	command: []string{"docker", "-v"},
	versionFunc: func(output string) string {
		/*
			Example:
			Docker version 23.0.1, build a5ee5b1dfc
		*/
		version, _, _ := strings.Cut(output[15:], ",")
		return "v" + version
	},
}
