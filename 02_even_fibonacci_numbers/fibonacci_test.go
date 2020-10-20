package main

import (
	"reflect"
	"testing"
)

func TestGetFibonacciTill(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  []int
	}{
		{
			"till 1",
			1,
			[]int{1},
		},
		{
			"till 2",
			2,
			[]int{1, 2},
		},
		{
			"till 3",
			3,
			[]int{1, 2, 3},
		},
		{
			"till 100",
			100,
			[]int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetFibonacciTill(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetFibonacciTill(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsEvenValue(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"1",
			1,
			false,
		},
		{
			"2",
			2,
			true,
		},
		{
			"100",
			100,
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsEvenValue(tc.input)
			if got != tc.want {
				t.Errorf("IsEvenValue(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
