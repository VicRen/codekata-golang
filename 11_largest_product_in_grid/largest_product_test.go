package main

import (
	"reflect"
	"testing"
)

func TestGetAdjacent(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		x     int
		y     int
		want  int
	}{
		{
			"out_of_range",
			[][]int{
				{1, 2},
				{1, 2},
			},
			0,
			3,
			0,
		},
		{
			"out_of_range_2",
			[][]int{
				{1, 2},
				{1, 2},
			},
			-1,
			1,
			0,
		},
		{
			"get_2",
			[][]int{
				{1, 2},
				{1, 2},
			},
			1,
			1,
			2,
		},
		{
			"get_1",
			[][]int{
				{1, 2},
				{1, 2},
			},
			0,
			0,
			1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAdjacent(tc.input, tc.x, tc.y)
			if got != tc.want {
				t.Errorf("GetAdjacent(%v, %d, %d)=%d, want %d", tc.input, tc.x, tc.y, got, tc.want)
			}
		})
	}
}

func TestGetRight4(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			"not_enough_digits",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			nil,
		},
		{
			"exact_four",
			[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
			[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
		},
		{
			"more_than_four_1",
			[][]int{
				{1, 2, 3, 4, 5},
			},
			[][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
			},
		},
		{
			"more_than_four_2",
			[][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			},
			[][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
				{1, 2, 3, 4},
				{2, 3, 4, 5},
			},
		},
		{
			"more_than_four_3",
			[][]int{
				{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
				{3, 4, 5, 6},
			},
		},
		{
			"more_than_four_4",
			[][]int{
				{1, 2, 3, 4, 5, 6},
				{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
				{3, 4, 5, 6},
				{1, 2, 3, 4},
				{2, 3, 4, 5},
				{3, 4, 5, 6},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetRight4(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetRight4(%v)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestGetDown4(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			"not_enough_lines",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			nil,
		},
		{
			"one_digits",
			[][]int{
				{1},
				{1},
				{1},
				{1},
			},
			[][]int{{1, 1, 1, 1}},
		},
		{
			"one_digits_2",
			[][]int{
				{1, 2},
				{1},
				{1, 3},
				{1},
			},
			[][]int{{1, 1, 1, 1}},
		},
		{
			"more_than_one_digits",
			[][]int{
				{1, 2},
				{1, 3},
				{1, 3},
				{1, 5},
			},
			[][]int{
				{1, 1, 1, 1},
				{2, 3, 3, 5},
			},
		},
		{
			"more_than_one_digits_2",
			[][]int{
				{1, 2, 3},
				{1, 3, 4},
				{1, 3},
				{1, 5},
			},
			[][]int{
				{1, 1, 1, 1},
				{2, 3, 3, 5},
			},
		},
		{
			"more_than_one_digits_3",
			[][]int{
				{1, 2, 3},
				{1, 3, 4},
				{1, 3, 5},
				{1, 5, 10},
			},
			[][]int{
				{1, 1, 1, 1},
				{2, 3, 3, 5},
				{3, 4, 5, 10},
			},
		},
		{
			"more_than_one_digits_4",
			[][]int{
				{1, 2},
				{1},
				{1, 3},
				{1},
				{2},
			},
			[][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 2},
			},
		},
		{
			"more_than_one_digits_5",
			[][]int{
				{1, 2},
				{1, 1},
				{1, 3},
				{1, 4},
				{2, 5},
			},
			[][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 2},
				{2, 1, 3, 4},
				{1, 3, 4, 5},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetDown4(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetDown4(%v)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestGetDiagonal4(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			"not_enough_lines",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			nil,
		},
		{
			"not_enough_columns",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			nil,
		},
		{
			"one_result",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3, 4},
			},
			[][]int{{1, 2, 3, 4}},
		},
		{
			"more_than_one_result",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3, 4},
				{1, 2, 3, 4, 5},
			},
			[][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
			},
		},
		{
			"more_than_one_result_2",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
			[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
		},
		{
			"more_than_one_result_3",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3, 4},
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4},
			},
			[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{2, 3, 4, 5},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetDiagonal4(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetDiagonal4(%v)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
