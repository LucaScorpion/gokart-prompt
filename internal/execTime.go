package internal

import (
	"fmt"
	"gokart-prompt/internal/ansi"
	"math"
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

	return ansi.Color(ansi.Yellow, " took "+formatMinutesSeconds(execTime))
}

func formatMinutesSeconds(totalSeconds float64) string {
	seconds := math.Mod(totalSeconds, 60)
	result := fmt.Sprintf("%.1fs", seconds)

	minutes := math.Floor(totalSeconds / 60)
	if minutes > 0 {
		result = fmt.Sprintf("%.0fm %s", minutes, result)
	}

	return result
}
