package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var dockerFiles = []string{
	"Dockerfile",
	"docker-compose.yml",
	"docker-compose.yaml",
}

func DockerVersion() string {
	if _, ok := internal.UpsearchWd(dockerFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("docker", "-v"); ok {
		/*
			Example:
			Docker version 23.0.1, build a5ee5b1dfc
		*/
		version, _, _ = strings.Cut(version[15:], ",")
		return ansi.Color(ansi.Blue, " 🐳 v"+version)
	}

	return ""
}
