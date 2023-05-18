package internal

import (
	"gokart-prompt/internal/ansi"
	"os"
)

func ExitCode() string {
	code := os.Getenv("EXIT_CODE")
	if code == "0" {
		return ansi.Color(ansi.Green, "➜ ")
	}
	return ansi.Color(ansi.Red, code+" ➜ ")
}
