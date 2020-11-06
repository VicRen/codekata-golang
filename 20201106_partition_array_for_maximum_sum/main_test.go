package main

import (
	"testing"
)

func Test_findMax(t *testing.T) {
	type args struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example_1",
			args{3, []int{1, 15, 7, 9, 2, 5, 10}},
			84,
		},
		{
			"example_2",
			args{4, []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}},
			83,
		},
		{
			"example_3",
			args{1, []int{1}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSum(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("findMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
