package main

import (
	"math/big"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fn, _ := (&big.Int{}).SetString("10203040506070809", 0)
	fn2, _ := (&big.Int{}).SetString("1234567890987654321", 0)
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{fn},
			false,
		},
		{
			"test2",
			args{fn2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.n); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLychrelNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{47},
			true,
		},
		{
			"test1",
			args{4994},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLychrelNumber(tt.args.n); got != tt.want {
				t.Errorf("isLychrelNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
