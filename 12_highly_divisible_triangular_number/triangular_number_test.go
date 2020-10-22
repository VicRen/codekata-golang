package main

import (
	"reflect"
	"testing"
)

func TestTau(t *testing.T) {
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
			"2",
			2,
			2,
		},
		{
			"10",
			10,
			4,
		},
		{
			"28",
			28,
			6,
		},
		{
			"76576500",
			76576500,
			576,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Tau(tc.input)
			if got != tc.want {
				t.Errorf("Tau(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestGetTriangularNumber(t *testing.T) {
	tt := []struct {
		name  string
		count int
		want  int
	}{
		{
			"1",
			1,
			1,
		},
		{
			"2",
			2,
			3,
		},
		{
			"7",
			7,
			28,
		},
		{
			"10",
			10,
			55,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GetTriangularNumber(tc.count)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetTriangularNumber(%d)=%v, want %v", tc.count, got, tc.want)
			}
		})
	}
}

func TestListDivisors(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  []int
	}{
		{
			"1",
			1,
			[]int{1},
		},
		{
			"3",
			3,
			[]int{1, 3},
		},
		{
			"6",
			6,
			[]int{1, 2, 3, 6},
		},
		{
			"28",
			28,
			[]int{1, 2, 4, 7, 14, 28},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ListDivisors(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ListDivisors(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
