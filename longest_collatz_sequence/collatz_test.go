package main

import "testing"

func TestCollatzNext(t *testing.T) {
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
			1,
		},
		{
			"3",
			3,
			10,
		},
		{
			"13",
			13,
			40,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := CollatzNext(tc.input)
			if got != tc.want {
				t.Errorf("CollatzNext(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
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
			"99",
			99,
			false,
		},
		{
			"100",
			100,
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsEven(tc.input)
			if got != tc.want {
				t.Errorf("IsEven(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
