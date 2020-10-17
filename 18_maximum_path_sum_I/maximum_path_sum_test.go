package main

import (
	"testing"
)

func TestLargestSumOfTriangle(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			"example",
			[][]int{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}},
			23,
		},
		{
			"example_2",
			[][]int{{3}, {7, 4}},
			10,
		},
		{
			"example_3",
			[][]int{{3}},
			3,
		},
		{
			"example_4",
			[][]int{{3}, {7, 4}, {2, 4, 6}},
			14,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := LargestSumOfTriangle(tc.input)
			if got != tc.want {
				t.Errorf("LargestSumOfTriangle(%v)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
