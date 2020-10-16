package main

import (
	"testing"
)

func TestCountNumberLetter(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{
			"1",
			1,
			3,
		},
		{
			"15",
			15,
			7,
		},
		{
			"90",
			90,
			6,
		},
		{
			"100",
			100,
			10,
		},
		{
			"300",
			300,
			12,
		},
		{
			"23",
			23,
			11,
		},
		{
			"123",
			123,
			24,
		},
		{
			"120",
			120,
			19,
		},
		{
			"115",
			115,
			20,
		},
		{
			"320",
			320,
			21,
		},
		{
			"444",
			444,
			23,
		},
		{
			"499",
			499,
			24,
		},
		{
			"342",
			342,
			23,
		},
		{
			"110",
			110,
			16,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := CountNumberLetter(tc.input)
			if got != tc.want {
				t.Errorf("CountNumberLetter(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
