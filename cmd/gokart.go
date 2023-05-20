package main

import (
	"fmt"
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"gokart-prompt/internal/git"
	"gokart-prompt/internal/versions"
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

	// Path and project info
	fmt.Print(internal.Path())
	fmt.Print(git.Git())

	// Languages
	fmt.Print(versions.GoVersion())
	fmt.Print(versions.JavaVersion())
	fmt.Print(versions.Node.GetVersion())
	fmt.Print(versions.PhpVersion())
	fmt.Print(versions.RubyVersion())
	fmt.Print(versions.RustVersion())

	// Tools
	fmt.Print(versions.Docker.GetVersion())

	fmt.Print(ansi.Reset())
}

func ps2() {
	fmt.Print(internal.ExitCode())
}
