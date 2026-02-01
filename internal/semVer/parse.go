package semVer

import (
	"strconv"
	"strings"
)

func Parse(v string) *SemVer {
	v = strings.TrimSpace(v)

	// Strip a leading 'v'.
	if len(v) > 0 && v[0] == 'v' {
		v = v[1:]
	}

	parts := strings.Split(v, ".")

	return &SemVer{
		major: parsePart(parts, 0),
		minor: parsePart(parts, 1),
		patch: parsePart(parts, 2),
	}
}

func parsePart(parts []string, index int) int {
	if index >= len(parts) {
		return 0
	}
	result, _ := strconv.Atoi(parts[index])
	return result
}
