package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"1",
			1,
			true,
		},
		{
			"11",
			11,
			true,
		},
		{
			"1001",
			1001,
			true,
		},
		{
			"123456654321",
			123456654321,
			true,
		},
		{
			"123456754321",
			123456754321,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("IsPalindrome(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func Test_findMaxPalindrome(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{10, 99},
			9009,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxPalindrome(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("findMaxPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
