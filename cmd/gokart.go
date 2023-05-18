package main

import (
	"fmt"
	"gokart-powerline/internal"
)

func main() {
	fmt.Print("\n")
	fmt.Print(internal.Path())

	fmt.Print("\n")
	fmt.Print(internal.ExitCode())
}
