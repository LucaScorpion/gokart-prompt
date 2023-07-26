package internal

import "strings"

func SemVerMatches(current, expected string) bool {
	expected = strings.TrimSpace(expected)
	if expected == "" {
		return true
	}

	// Strip a leading 'v' from both versions.
	if current[0] == 'v' {
		current = current[1:]
	}
	if expected[0] == 'v' {
		expected = expected[1:]
	}

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
