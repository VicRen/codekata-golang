package main

import "testing"

func Test_isDivisible(t *testing.T) {
	tt := []struct {
		name  string
		start int
		end   int
		input int
		want  bool
	}{
		{
			"2",
			1,
			2,
			2,
			true,
		},
		{
			"10",
			1,
			10,
			2520,
			true,
		},
		{
			"20",
			1,
			20,
			20,
			false,
		},
		{
			"14",
			1,
			20,
			14,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsDivisible(tc.start, tc.end, tc.input)
			if got != tc.want {
				t.Errorf("IsDivisible(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
