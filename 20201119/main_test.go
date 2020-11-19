package main

import "testing"

func Test_multiple(t *testing.T) {
	type args struct {
		n uint64
		m uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"example_1",
			args{1, 10},
			10,
		},
		{
			"example_2",
			args{3, 4},
			12,
		},
		{
			"example_3",
			args{100, 100},
			10000,
		},
		{
			"example_4",
			args{300, 100},
			30000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiple(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiple2(t *testing.T) {
	type args struct {
		n uint64
		m uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"example_1",
			args{1, 10},
			10,
		},
		{
			"example_2",
			args{3, 4},
			12,
		},
		{
			"example_3",
			args{100, 100},
			10000,
		},
		{
			"example_4",
			args{300, 100},
			30000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiple2(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("multiple2() = %v, want %v", got, tt.want)
			}
		})
	}
}
