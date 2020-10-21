package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"1",
			1,
			true,
		},
		{
			"11",
			11,
			true,
		},
		{
			"1001",
			1001,
			true,
		},
		{
			"123456654321",
			123456654321,
			true,
		},
		{
			"123456754321",
			123456754321,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("IsPalindrome(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
