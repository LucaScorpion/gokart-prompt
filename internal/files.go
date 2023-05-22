package internal

import (
	"os"
)

func ListWdFiles() []string {
	wd, err := os.Getwd()
	if err != nil {
		return nil
	}

	entries, err := os.ReadDir(wd)
	if err != nil {
		return nil
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files
}
