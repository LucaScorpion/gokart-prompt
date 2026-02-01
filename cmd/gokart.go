package main

import (
	"fmt"
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"gokart-prompt/internal/git"
	"gokart-prompt/internal/terminal"
	"gokart-prompt/internal/versions"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Expected at most one argument: ps1 or ps2")
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		info()
		return
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

func info() {
	fmt.Println("Gokart Prompt 🏎")
	fmt.Println("For setup, source code, and info, see: https://github.com/LucaScorpion/gokart-prompt")
}

func ps1() {
	wdFiles := internal.ListWdFiles()

	fmt.Println(terminal.RightAlign(internal.Time()))

	fmt.Print(ansi.Bold())
	fmt.Print(internal.Path())
	fmt.Print(git.Git())
	fmt.Print(versions.All(wdFiles))
	fmt.Print(internal.CmdTime())
	fmt.Println(ansi.Reset())

	ps2()
}

func ps2() {
	fmt.Print(internal.ExitCode())
}
