package semVer

import (
	"fmt"
	"strings"
)

type SemVer struct {
	major int
	minor int
	patch int
}

func (s SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}

func normaliseSemVer(v string) string {
	v = strings.TrimSpace(v)
	// Strip a leading 'v'.
	if len(v) > 0 && v[0] == 'v' {
		v = v[1:]
	}
	return v
}
