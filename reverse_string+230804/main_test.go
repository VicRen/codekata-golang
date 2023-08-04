package main

import "testing"

func ReverseString(s string) string {
	tmp := []byte(s)
	l := len(tmp)
	for i := 0; i < l/2; i++ {
		tmp[i], tmp[l-i-1] = tmp[l-i-1], tmp[i]
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
	}

	for _, tc := range tt {
		got := ReverseString(tc.input)
		if got != tc.want {
			t.Errorf("ReverseString(%s) = %s, want %s", tc.input, got, tc.want)
		}
	}
}
