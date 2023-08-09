package main

import "testing"

func ReverseString(s string) string {
	ret := []rune(s)
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}
	return string(ret)
}

func TestReverseString(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{"test", "test", "tset"},
		{"testing", "testing", "gnitset"},
		{"i am testing", "i am testing", "gnitset ma i"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseString(tc.input)
			if got != tc.want {
				t.Errorf("ReverseString(%s) = %s, want %s\n", tc.input, got, tc.want)
			}
		})
	}
}
