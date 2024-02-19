package internal

import (
	"gokart-prompt/internal/ansi"
	"os"
	"testing"
)

func TestExitCode(t *testing.T) {
	cases := map[string]string{
		"0":   ansi.Color(ansi.Green, "➜ "),
		"1":   ansi.Color(ansi.Red, "1 ➜ "),
		"130": ansi.Color(ansi.Red, "130 ➜ "),
	}

	for exitCode, expected := range cases {
		os.Setenv("EXIT_CODE", exitCode)
		actual := ExitCode()
		if actual != expected {
			t.Errorf("EXIT_CODE=%s ExitCode() = %s, expected %s", exitCode, actual, expected)
		}
	}
}
