package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactorsOf(t *testing.T) {
	tt := []struct{
		name string
		input int
		want []int
	}{
		{
			"1",
			1,
			[]int{},
		},
		{
			"2",
			2,
			[]int{2},
		},
		{
			"3",
			3,
			[]int{3},
		},
		{
			"4",
			4,
			[]int{2, 2},
		},
		{
			"5",
			5,
			[]int{5},
		},
		{
			"6",
			6,
			[]int{2, 3},
		},
		{
			"7",
			7,
			[]int{7},
		},
		{
			"8",
			8,
			[]int{2, 2, 2},
		},
		{
			"9",
			9,
			[]int{3, 3},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimeFactorsOf(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PrimeFactorsOf %d = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}