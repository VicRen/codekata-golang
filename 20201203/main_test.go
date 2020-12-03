package main

import "testing"

func Test_goodPairsCount(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{1, 2, 3, 1, 1, 3}},
			4,
		},
		{
			"test2",
			args{[]int{1, 1, 1, 1}},
			6,
		},
		{
			"test3",
			args{[]int{1, 2, 3}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodPairsCount(tt.args.src); got != tt.want {
				t.Errorf("goodPairsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goodPairsCount2(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{1, 2, 3, 1, 1, 3}},
			4,
		},
		{
			"test2",
			args{[]int{1, 1, 1, 1}},
			6,
		},
		{
			"test3",
			args{[]int{1, 2, 3}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodPairsCount2(tt.args.src); got != tt.want {
				t.Errorf("goodPairsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
