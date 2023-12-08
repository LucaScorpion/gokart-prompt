package internal

import "testing"

func TestFormatMinutesSeconds(t *testing.T) {
	cases := map[float64]string{
		3.14:   "3.1s",
		60:     "1m 0.0s",
		3712.5: "61m 52.5s",
	}

	for in, expected := range cases {
		actual := formatMinutesSeconds(in)
		if actual != expected {
			t.Errorf("formatMinutesSeconds(%f) = %s, expected %s", in, actual, expected)
		}
	}
}
