package main

import (
	"testing"
)

func Test_findFactorial(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  string
	}{
		{
			"2",
			2,
			"2",
		},
		{
			"4",
			4,
			"24",
		},
		{
			"10",
			10,
			"3628800",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := findFactorial(tc.input)
			if got != tc.want {
				t.Errorf("findFactorial=%v, want %v", got, tc.want)
			}
		})
	}
}

func Test_sumDigits(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  int
	}{
		{
			"no_text",
			"",
			0,
		},
		{
			"1",
			"1",
			1,
		},
		{
			"12345",
			"12345",
			15,
		},
		{
			"3628800",
			"3628800",
			27,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := sumDigits(tc.input)
			if got != tc.want {
				t.Errorf("sumDigits = %d, want %d", got, tc.want)
			}
		})
	}
}
