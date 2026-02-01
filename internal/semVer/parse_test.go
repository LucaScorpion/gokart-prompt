package semVer

import "testing"

func TestParse(t *testing.T) {
	cases := [][]any{
		{"1.2.3", SemVer{1, 2, 3, ""}},
		{"v1.2.3", SemVer{1, 2, 3, ""}},
		{"1.2", SemVer{1, 2, 0, ""}},
		{"v1.2", SemVer{1, 2, 0, ""}},
		{"1", SemVer{1, 0, 0, ""}},
		{"v1", SemVer{1, 0, 0, ""}},
		{"  1.2.3  ", SemVer{1, 2, 3, ""}},
		{"  v1.2.3  ", SemVer{1, 2, 3, ""}},
		{"invalid", SemVer{0, 0, 0, ""}},
		{"5.6.what", SemVer{5, 6, 0, ""}},
	}

	for _, c := range cases {
		input := c[0].(string)
		expected := c[1].(SemVer)
		actual := *Parse(input)
		if actual != expected {
			t.Errorf("Parse(%s) = %v, expected %v", input, actual, expected)
		}
	}
}
