package main

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{
			"largest prime factor of 1",
			1,
			1,
		},
		{
			"largest prime factor of 2",
			2,
			2,
		},
		{
			"largest prime factor of 4",
			4,
			2,
		},
		{
			"largest prime factor of a large number",
			2 * 2 * 2 * 3 * 3 * 5 * 17 * 37,
			37,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := LargestPrimeFactor(tc.input)
			if got != tc.want {
				t.Errorf("LargestPrimeFactor(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
