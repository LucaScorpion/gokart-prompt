package versions

import (
	"gokart-powerline/internal"
	"gokart-powerline/internal/ansi"
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
		version, _, _ = strings.Cut(version, ",")
		parts := strings.Split(version, " ")
		version = parts[len(parts)-1]

		return ansi.Color(ansi.Blue, " 🐳 v"+version)
	}

	return ""
}
