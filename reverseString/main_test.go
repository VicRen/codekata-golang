package main

import "testing"

func ReverseString(s string) string {
	tm := []byte(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		tm[i], tm[l-1-i] = tm[l-1-i], tm[i]
	}
	return string(tm)
}

func TestReverseString(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{"1", "test", "tset"},
		{"2", "testing", "gnitset"},
		{"3", "this is a testing string", "gnirts gnitset a si siht"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseString(tc.input)
			if got != tc.want {
				t.Errorf("ReverseString(%s) = %s, want %s", tc.input, got, tc.want)
			}
		})
	}
}
