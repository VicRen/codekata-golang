package main

import (
	"reflect"
	"testing"
)

func Test_isAbundantNumber(t *testing.T) {
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
			"12",
			12,
			true,
		},
		{
			"28",
			28,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := isAbundantNumber(tc.input)
			if got != tc.want {
				t.Errorf("isAbundantNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_sumOfDivisors(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{
			"1",
			1,
			0,
		},
		{
			"2",
			2,
			1,
		},
		{
			"4",
			4,
			3,
		},
		{
			"12",
			12,
			16,
		},
		{
			"28",
			28,
			28,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := sumOfDivisors(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("sumOfDivisors() = %d, want %d", got, tc.want)
			}
		})
	}
}
