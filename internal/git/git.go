package git

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
)

const branchIcon = "\uE0A0" // 

func Git() string {
	if rawStatus, ok := internal.Command("git", "status", "--porcelain", "-b"); ok {
		s := parseStatus(rawStatus)

		return ansi.Color(ansi.Magenta, " "+branchIcon+" "+s.branch)
	}

	return ""
}
