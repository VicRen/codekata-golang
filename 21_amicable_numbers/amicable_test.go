package main

import (
	"testing"
)

func Test_sumOfDivisors(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{
			"220",
			220,
			284,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := sumOfDivisors(tc.input)
			if got != tc.want {
				t.Errorf("sumOfDivisors() = %d, want %d", got, tc.want)
			}
		})
	}
}

func Test_isAmicableNumber(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"220",
			220,
			true,
		},
		{
			"221",
			221,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := isAmicableNumber(tc.input)
			if got != tc.want {
				t.Errorf("isAmicableNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}
