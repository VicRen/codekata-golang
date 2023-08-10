package main

import "testing"

func ReverseString(s string) string {
	r := []rune(s)
	for l, e := 0, len(r)-1; l < e; l, e = l+1, e-1 {
		r[l], r[e] = r[e], r[l]
	}
	return string(r)
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
