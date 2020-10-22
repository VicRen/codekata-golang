package main

import "testing"

func TestIsPrime(t *testing.T) {
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
			"10",
			10,
			false,
		},
		{
			"17",
			17,
			true,
		},
		{
			"97",
			97,
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPrime(tc.input)
			if got != tc.want {
				t.Errorf("IsPrime(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
