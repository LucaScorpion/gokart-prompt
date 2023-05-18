package versions

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"strings"
)

var javaFiles = []string{
	"pom.xml",
	"build.gradle",
	"settings.gradle",
	"build.xml",
}

func JavaVersion() string {
	if _, ok := internal.UpsearchWd(javaFiles); !ok {
		return ""
	}

	if version, ok := internal.Command("java", "--version"); ok {
		/*
			Example:
			openjdk 19.0.2 2023-01-17
			OpenJDK Runtime Environment (build 19.0.2+7)
			OpenJDK 64-Bit Server VM (build 19.0.2+7, mixed mode)
		*/
		version := strings.SplitN(version, " ", 3)[1]
		return ansi.Color(ansi.Cyan, " ☕ v"+version)
	}

	return ""
}
