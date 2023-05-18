package git

import "strings"

type status struct {
	branch    string
	untracked bool
	staged    bool
	unstaged  bool
	renamed   bool
	deleted   bool
	unmerged  bool
}

func parseStatus(raw string) status {
	lines := strings.Split(raw, "\n")
	result := status{}

	for _, line := range lines {
		parseStatusLine(&result, line)
	}

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
