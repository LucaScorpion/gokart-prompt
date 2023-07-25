package internal

import (
	"fmt"
	"gokart-prompt/internal/ansi"
	"os"
	"strconv"
)

func CmdTime() string {
	// If no command was entered, never display the time.
	if os.Getenv("GOKART_CMD") == "" {
		return ""
	}

	// Try to parse the command start and end times.
	cmdStart, err := strconv.ParseFloat(os.Getenv("GOKART_CMD_START"), 64)
	if err != nil {
		return ""
	}
	cmdEnd, err := strconv.ParseFloat(os.Getenv("GOKART_CMD_END"), 64)
	if err != nil {
		return ""
	}

	execTime := cmdEnd - cmdStart
	if execTime < 1 {
		return ""
	}

	return ansi.Color(ansi.Yellow, " took "+fmt.Sprintf("%.1f", execTime)+"s")
}
