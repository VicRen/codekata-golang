package main

import "testing"

func TestIsDivisibleFor1to20(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"20",
			20,
			false,
		},
		{
			"14",
			14,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsDivisibleFor1to20(tc.input)
			if got != tc.want {
				t.Errorf("IsDivisibleFor1to20(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
