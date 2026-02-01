package semVer

import "strings"

func Matches(current, expected string) bool {
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
