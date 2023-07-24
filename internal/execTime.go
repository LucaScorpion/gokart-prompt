package internal

import (
	"fmt"
	"gokart-prompt/internal/ansi"
	"os"
	"strconv"
)

func CmdTime(now int64) string {
	preCmdTime, err := strconv.ParseFloat(os.Getenv("GOKART_TIME"), 64)
	if err != nil {
		return ""
	}

	execTime := float64(now)/1000 - preCmdTime
	return ansi.Color(ansi.Yellow, " took "+fmt.Sprintf("%.1f", execTime)+"s")
}
