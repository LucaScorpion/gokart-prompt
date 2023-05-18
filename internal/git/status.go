package git

import "strings"

type status struct {
	branch string
}

func parseStatus(raw string) status {
	lines := strings.Split(raw, "\n")
	result := status{}

	for _, line := range lines {
		parseStatusLine(&result, line)
	}

	return result
}

func parseStatusLine(s *status, line string) {
	code := strings.TrimSpace(line[:2])

	switch code {
	case "##":
		branch, _, _ := strings.Cut(line[3:], "...")
		s.branch = branch
	}
}
