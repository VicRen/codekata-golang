package main

import "testing"

func ReverseString(s string) string {
	tmp := []rune(s)
	l := len(tmp)
	for i := 0; i < l/2; i++ {
		tmp[i], tmp[l-i-1] = tmp[l-i-1], tmp[i]
	}
	return string(tmp)
}

func ReverseString2(s string) string {
	tmp := []rune(s)
	for f, t := 0, len(tmp)-1; t > f; f, t = f+1, t-1 {
		tmp[f], tmp[t] = tmp[t], tmp[f]
	}
	return string(tmp)
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
				t.Errorf("ReverseString(%s) = %s, want %s", tc.input, got, tc.want)
			}
			got = ReverseString2(tc.input)
			if got != tc.want {
				t.Errorf("ReverseString(%s) = %s, want %s", tc.input, got, tc.want)
			}
		})
	}
}
