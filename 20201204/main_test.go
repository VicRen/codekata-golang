package main

import (
	"testing"
)

func Test_countMove(t *testing.T) {
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
			args{[]int{1, 2, 3}},
			3,
		},
		{
			"test2",
			args{[]int{1, 1, 1}},
			0,
		},
		{
			"test3",
			args{[]int{-100, 0, 100}},
			300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMove(tt.args.src); got != tt.want {
				t.Errorf("countMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
