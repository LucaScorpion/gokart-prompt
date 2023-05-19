package git

import (
	"gokart-prompt/internal"
	"strconv"
	"strings"
)

type status struct {
	branch string

	untracked bool
	staged    bool
	unstaged  bool
	renamed   bool
	deleted   bool
	unmerged  bool

	ahead  bool
	behind bool
}

func parseStatus(raw string) status {
	lines := strings.Split(raw, "\n")
	result := status{}

	for _, line := range lines {
		parseStatusLine(&result, line)
	}

	aheadCount := revListCount(result.branch + "@{upstream}..HEAD")

	// If there is no upstream branch, revListCount will return -1.
	// This essentially means we are "ahead", and should display the push icon.
	if aheadCount == -1 {
		aheadCount = 1
	}

	result.ahead = aheadCount != 0
	result.behind = revListCount("HEAD.."+result.branch+"@{upstream}") > 0

	return result
}

// See: https://git-scm.com/docs/git-status#_output
func parseStatusLine(s *status, line string) {
	code := line[:2]

	if code == "##" {
		branch, _, _ := strings.Cut(line[3:], "...")
		s.branch = branch
		return
	}
	if code == "??" {
		s.untracked = true
		return
	}

	codeX := rune(code[0]) // Index
	codeY := rune(code[1]) // Working tree

	if codeX == 'M' || codeX == 'T' || codeX == 'A' {
		s.staged = true
	}

	if codeY == 'M' || codeY == 'T' {
		s.unstaged = true
	}

	if codeX == 'R' || codeY == 'R' {
		s.renamed = true
	}

	if codeX == 'D' || codeY == 'D' {
		s.deleted = true
	}

	if codeX == 'U' || codeY == 'U' || code == "DD" || code == "AA" {
		s.unmerged = true
	}
}

func revListCount(commit string) int {
	c, _ := internal.Command("git", "rev-list", "--count", commit)

	if i, err := strconv.Atoi(c); err == nil {
		return i
	}

	return -1
}
