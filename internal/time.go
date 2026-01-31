package internal

import (
	"time"
)

func Time() string {
	return time.Now().Format(time.TimeOnly)
}
