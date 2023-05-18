package internal

import (
	"os"
)

func ExitCode() string {
	code := os.Getenv("EXIT_CODE")
	if code == "0" {
		return ""
	}
	return code + " "
}
