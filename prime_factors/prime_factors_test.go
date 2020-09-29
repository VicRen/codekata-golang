package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactorsOf(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  []int
	}{
		{
			"prime factors of 1",
			1,
			[]int{},
		},
		{
			"prime factors of 2",
			2,
			[]int{2},
		},
		{
			"prime factors of 3",
			3,
			[]int{3},
		},
		{
			"prime factors of 1",
			1,
			[]int{},
		},
		{
			"prime factors of 2",
			2,
			[]int{2},
		},
		{
			"prime factors of 4",
			4,
			[]int{2, 2},
		},
		{
			"prime factors of 5",
			5,
			[]int{5},
		},
		{
			"prime factors of 6",
			6,
			[]int{2, 3},
		},
		{
			"prime factors of 7",
			7,
			[]int{7},
		},
		{
			"prime factors of 8",
			8,
			[]int{2, 2, 2},
		},
		{
			"prime factors of 9",
			9,
			[]int{3, 3},
		},
		{
			"prime factors of a big number",
			2 * 2 * 2 * 3 * 3 * 5 * 7 * 11 * 13 * 73,
			[]int{2, 2, 2, 3, 3, 5, 7, 11, 13, 73},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimeFactorsOf(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PrimeFactorsOf %d, got %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
