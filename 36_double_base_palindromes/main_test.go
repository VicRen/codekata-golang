package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
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
			got := isPalindromeDecimal(tc.input)
			if got != tc.want {
				t.Errorf("IsPalindrome(%d)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func Test_isPalindromeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"1"},
			true,
		},
		{
			"11",
			args{"11"},
			true,
		},
		{
			"1001",
			args{"1001"},
			true,
		},
		{
			"123456654321",
			args{"123456654321"},
			true,
		},
		{
			"123456754321",
			args{"123456754321"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToBinary(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{1},
			"1",
		},
		{
			"2",
			args{2},
			"10",
		},
		{
			"3",
			args{3},
			"11",
		},
		{
			"7",
			args{7},
			"111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBinary(tt.args.num); got != tt.want {
				t.Errorf("convertToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
