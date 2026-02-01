package semVer

import (
	"fmt"
)

type SemVer struct {
	major    int
	minor    int
	patch    int
	operator string
}

type Operator string

var Tilde Operator = "~"
var Caret Operator = "^"

func (s SemVer) String() string {
	return fmt.Sprintf("%s%d.%d.%d", s.operator, s.major, s.minor, s.patch)
}
