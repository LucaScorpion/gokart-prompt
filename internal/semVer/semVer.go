package semVer

import (
	"fmt"
)

type SemVer struct {
	major int
	minor int
	patch int
}

type SemVerRange struct {
	SemVer
	operator string
}

type Operator string

var Gt Operator = ">"
var Gte Operator = ">="
var Lt Operator = "<"
var Lte Operator = "<="

func (s SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}
