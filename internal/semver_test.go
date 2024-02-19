package internal

import "testing"

func TestSemVerMatches(t *testing.T) {
	cases := [][]any{
		{"1.2.3", "1.2.3", true},
		{"1.2.3", "1.2", true},
		{"1.2.3", "1", true},
		{"invalid", "2.0.0", false},
		{"2.0.0", "invalid", false},
		{"", "3.0.0", false},
		{"4.0.0", "v4", true},
		{"v4.0.0", "4", true},
		{" 5.0.0 ", "5.0.0", true},
		{"5.0.0 ", "  5.0.0  ", true},
	}

	for _, c := range cases {
		a := c[0].(string)
		b := c[1].(string)
		expected := c[2].(bool)
		actual := SemVerMatches(a, b)
		if actual != expected {
			t.Errorf("SemVerMatches(%s, %s) = %t, expected %t", a, b, actual, expected)
		}
	}
}
