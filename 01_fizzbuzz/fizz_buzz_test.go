package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  string
	}{
		{
			"case 1",
			1,
			"",
		},
		{
			"case 3",
			3,
			"Fizz",
		},
		{
			"case 5",
			5,
			"Buzz",
		},
		{
			"case 6",
			6,
			"Fizz",
		},
		{
			"case 10",
			10,
			"Buzz",
		},
		{
			"case 15",
			15,
			"FizzBuzz",
		},
		{
			"case 45",
			45,
			"FizzBuzz",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := FizzBuzz(tc.input)
			if got != tc.want {
				t.Errorf("FizzBuzz(%d)=%s, want %s", tc.input, got, tc.want)
			}
		})
	}
}
