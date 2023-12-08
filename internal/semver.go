package internal

import "strings"

func SemVerMatches(current, expected string) bool {
	expected = normaliseSemVer(expected)
	if expected == "" {
		return true
	}
	current = normaliseSemVer(current)

	fullVersionParts := strings.Split(current, ".")
	checkParts := strings.Split(expected, ".")

	// Check if all the parts of the expected version match the current version.
	for i, part := range checkParts {
		if i >= len(fullVersionParts) || part != fullVersionParts[i] {
			return false
		}
	}

	return true
}

func normaliseSemVer(v string) string {
	v = strings.TrimSpace(v)
	// Strip a leading 'v'.
	if len(v) > 0 && v[0] == 'v' {
		v = v[1:]
	}
	return v
}
