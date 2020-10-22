package main

import (
	"reflect"
	"testing"
)

func TestSplitInteger(t *testing.T) {
	tt := []struct {
		name string
		src  int
		want [][]int
	}{
		{
			"5-3",
			6,
			[][]int{
				{1, 2, 3},
			},
		},
		{
			"10-3",
			10,
			[][]int{
				{1, 2, 7},
				{1, 3, 6},
				{1, 4, 5},
				{2, 3, 5},
			},
		},
		{
			"11-3",
			11,
			[][]int{
				{1, 2, 8},
				{1, 3, 7},
				{1, 4, 6},
				{2, 3, 6},
				{2, 4, 5},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitIntegerIntoThree(tc.src)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SplitInteger(%d)=%v, want %v", tc.src, got, tc.want)
			}
		})
	}
}

func TestIsPythagoreanTriplet(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			"3,4",
			[]int{3, 4},
			false,
		},
		{
			"3,4,5,6",
			[]int{3, 4, 5, 6},
			false,
		},
		{
			"3,4,5",
			[]int{3, 4, 5},
			true,
		},
		{
			"3,4,6",
			[]int{3, 4, 6},
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPythagoreanTriplet(tc.input)
			if got != tc.want {
				t.Errorf("IsPythagoreanTriplet(%v)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"0",
			[]int{0, 1, 1},
			0,
		},
		{
			"2",
			[]int{2, 1, 1},
			2,
		},
		{
			"3",
			[]int{2, 3, 4},
			24,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetProduct(tc.input)
			if got != tc.want {
				t.Errorf("GetProduct(%v)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
