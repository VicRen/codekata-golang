package main

import "testing"

func Test_findNumber(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"1",
			args{1},
			1,
		},
		{
			"2",
			args{2},
			3,
		},
		{
			"4",
			args{4},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumber(tt.args.n); got != tt.want {
				t.Errorf("findNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxLine(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"5",
			args{5},
			2,
		},
		{
			"8",
			args{8},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLine(tt.args.n); got != tt.want {
				t.Errorf("findMaxLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
