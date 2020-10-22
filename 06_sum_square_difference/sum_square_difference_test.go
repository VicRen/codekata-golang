package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{
			"1",
			1,
			1,
		},
		{
			"10",
			10,
			385,
		},
		{
			"5",
			5,
			1 + 4 + 9 + 16 + 25,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfSquares(tc.input)
			if got != tc.want {
				t.Errorf("SumOfSquares(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestSquareSum(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{
			"1",
			1,
			1,
		},
		{
			"5",
			5,
			15 * 15,
		},
		{
			"10",
			10,
			3025,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SquareSum(tc.input)
			if got != tc.want {
				t.Errorf("SquareSum(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
