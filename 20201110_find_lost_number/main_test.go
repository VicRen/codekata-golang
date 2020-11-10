package main

import (
	"testing"
)

func Test_findLostNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example_1",
			args{[]int{3, 0, 1}},
			2,
		},
		{
			"example_2",
			args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			8,
		},
		{
			"example_3",
			args{[]int{8, 6, 4, 2, 3, 5, 7, 0, 1}},
			9,
		},
		{
			"example_4",
			args{[]int{9, 8, 6, 4, 2, 3, 5, 7, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLostNumber(tt.args.nums); got != tt.want {
				t.Errorf("findLostNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLostNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example_1",
			args{[]int{3, 0, 1}},
			2,
		},
		{
			"example_2",
			args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			8,
		},
		{
			"example_3",
			args{[]int{9, 8, 6, 4, 2, 3, 5, 7, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLostNumber2(tt.args.nums); got != tt.want {
				t.Errorf("findLostNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}
