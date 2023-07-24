package main

import (
	"fmt"
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"gokart-prompt/internal/git"
	"gokart-prompt/internal/versions"
	"os"
	"time"
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
	nowMillis := time.Now().UnixMilli()
	wdFiles := internal.ListWdFiles()

	fmt.Print("\n")
	fmt.Print(ansi.Bold())
	fmt.Print(internal.Path())
	fmt.Print(git.Git())
	fmt.Print(versions.All(wdFiles))
	fmt.Print(internal.CmdTime(nowMillis))

	fmt.Print(ansi.Reset())
	fmt.Println()
	ps2()
}

func ps2() {
	fmt.Print(internal.ExitCode())
}
