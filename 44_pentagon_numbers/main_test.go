package main

import "testing"

func Test_pentagonalNumberN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{1},
			1,
		},
		{
			"2",
			args{2},
			5,
		},
		{
			"10",
			args{10},
			145,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pentagonalNumberN(tt.args.n); got != tt.want {
				t.Errorf("pentagonalNumberN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPentagonalNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{1},
			true,
		},
		{
			"5",
			args{5},
			true,
		},
		{
			"6",
			args{6},
			false,
		},
		{
			"120",
			args{120},
			false,
		},
		{
			"145",
			args{145},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPentagonalNumber(tt.args.n); got != tt.want {
				t.Errorf("isPentagonalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
