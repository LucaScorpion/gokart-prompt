package main

import (
	"fmt"
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"gokart-prompt/internal/git"
	"gokart-prompt/internal/terminal"
	"gokart-prompt/internal/versions"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected one argument: ps1 or ps2")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "ps1" {
		ps1()
	} else if arg == "ps2" {
		ps2()
	} else {
		fmt.Println("Argument should be ps1 or ps2")
		os.Exit(1)
	}
}

func ps1() {
	wdFiles := internal.ListWdFiles()

	fmt.Print("\n")

	// Left part of PS1.
	line := ansi.Bold() + internal.Path() + git.Git() + versions.All(wdFiles) + internal.CmdTime()
	fmt.Print(line)

	// Reset, spacing, and right part of PS1.
	fmt.Print(ansi.Reset())
	fmt.Print(rightAlign(line, " "+internal.Time()))

	// PS2 on a new line.
	fmt.Print(ansi.Reset())
	fmt.Println()
	ps2()
}

func ps2() {
	fmt.Print(internal.ExitCode())
}

func rightAlign(left, right string) string {
	columns := terminal.Columns()
	if columns <= 0 {
		return ""
	}

	leftLen := utf8.RuneCountInString(ansi.ToPlain(left))
	rightLen := utf8.RuneCountInString(ansi.ToPlain(right))
	return strings.Repeat(" ", terminal.Columns()-leftLen-rightLen-1) + right
}
