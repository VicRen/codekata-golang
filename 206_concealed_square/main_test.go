package main

import (
	"math/big"
	"testing"
)

func Test_isMatch(t *testing.T) {
	fn, _ := (&big.Int{}).SetString("10203040506070809", 0)
	fn2, _ := (&big.Int{}).SetString("10203040506070808", 0)
	type args struct {
		i *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				fn,
			},
			true,
		},
		{
			"test2",
			args{
				fn2,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.i); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEndsWith37(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				"123",
			},
			true,
		},
		{
			"test2",
			args{
				"127",
			},
			true,
		},
		{
			"test3",
			args{
				"12",
			},
			false,
		},
		{
			"test4",
			args{
				"",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEndsWith37(tt.args.s); got != tt.want {
				t.Errorf("isEndsWith37() = %v, want %v", got, tt.want)
			}
		})
	}
}
