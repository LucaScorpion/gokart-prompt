package internal

import (
	"os"
	"path"
)

func UpsearchWd(fileNames []string) (string, bool) {
	if wd, err := os.Getwd(); err != nil {
		return "", false
	} else {
		return upsearch(wd, fileNames)
	}
}

func upsearch(dir string, fileNames []string) (string, bool) {
	for _, name := range fileNames {
		p := path.Join(dir, name)
		if _, err := os.Stat(p); err == nil {
			return p, true
		}
	}

	parentDir := path.Dir(dir)
	if parentDir != dir && !isProjectDir(dir) {
		return upsearch(parentDir, fileNames)
	}

	return "", false
}
