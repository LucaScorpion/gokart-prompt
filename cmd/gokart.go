package main

import (
	"fmt"
	"gokart-powerline/internal"
	"gokart-powerline/internal/ansi"
	"gokart-powerline/internal/versions"
	"os"
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
	fmt.Print("\n")
	fmt.Print(ansi.Bold())

	fmt.Print(internal.Path())

	fmt.Print(versions.NodeVersion())

	fmt.Print(ansi.Reset())
}

func ps2() {
	fmt.Print(internal.ExitCode())
}
