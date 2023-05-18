package internal

import (
	"os"
	"path"
)

const gitFolder = ".git"

func IsProjectDir(dir string) bool {
	if f, err := os.Stat(path.Join(dir, gitFolder)); err != nil {
		return false
	} else {
		return f.IsDir()
	}
}
