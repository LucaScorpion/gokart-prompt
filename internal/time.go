package internal

import (
	"gokart-prompt/internal/ansi"
	"time"
)

func Time() string {
	return ansi.Color256(244, time.Now().Format(time.TimeOnly))
}
