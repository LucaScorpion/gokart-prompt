package git

import (
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
)

const branchIcon = "\uE0A0" // 

func Git() string {
	if rawStatus, ok := internal.Command("git", "status", "--porcelain", "-b"); ok {
		s := parseStatus(rawStatus)

		return ansi.Color(ansi.Magenta, " "+branchIcon+" "+s.branch) + statusText(s)
	}

	return ""
}

func statusText(s status) string {
	symbols := ""

	if s.untracked {
		symbols += "?"
	}
	if s.staged {
		symbols += "+"
	}
	if s.unstaged {
		symbols += "!"
	}
	if s.renamed {
		symbols += "»"
	}
	if s.deleted {
		symbols += "✘"
	}
	if s.unmerged {
		symbols += "="
	}

	if symbols != "" {
		symbols = ansi.Color(ansi.Red, " ["+symbols+"]")
	}

	return symbols
}
