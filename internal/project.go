package internal

import (
	"os"
	"path"
)

const gitFolder = ".git"

func isProjectDir(dir string) bool {
	if f, err := os.Stat(path.Join(dir, gitFolder)); err != nil {
		return false
	} else {
		return f.IsDir()
	}
}

func upsearchProjectDir() (string, bool) {
	name, isProject := UpsearchWd([]string{gitFolder})
	if isProject {
		name = path.Dir(name)
	}
	return name, isProject
}
