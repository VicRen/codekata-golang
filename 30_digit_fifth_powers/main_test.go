package main

import (
	"testing"
)

func Test_fifthPower(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"1",
			args{1},
			1,
		},
		{
			"2",
			args{2},
			32,
		},
		{
			"9",
			args{9},
			59049,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fifthPower(tt.args.n); got != tt.want {
				t.Errorf("fifthPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfFifthPowerOfDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"1",
			args{1},
			1,
		},
		{
			"11",
			args{11},
			2,
		},
		{
			"111",
			args{111},
			3,
		},
		{
			"222",
			args{222},
			96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfFifthPowerOfDigits(tt.args.n); got != tt.want {
				t.Errorf("sumOfFifthPowerOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
