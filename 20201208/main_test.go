package main

import "testing"

func Test_canSeparate(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				[]int{1, 1, 1},
			},
			true,
		},
		{
			"test1",
			args{
				[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			},
			true,
		},
		{
			"test2",
			args{
				[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			},
			false,
		},
		{
			"test3",
			args{
				[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			},
			true,
		},
		{
			"test4",
			args{
				[]int{10, -10, 10, -10, 10, -10, 10, -10},
			},
			false,
		},
		{
			"test5",
			args{
				[]int{1, -1, 1, -1},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSeparate(tt.args.src); got != tt.want {
				t.Errorf("canSeparate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumSlice(t *testing.T) {
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
			args{
				[]int{1, 2},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSlice(tt.args.src); got != tt.want {
				t.Errorf("sumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
