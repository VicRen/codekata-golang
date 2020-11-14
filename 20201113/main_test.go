package main

import "testing"

func Test_isCompliance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"example_1",
			args{[]int{4, 2, 3}},
			true,
		},
		{
			"example_2",
			args{[]int{4, 2, 1}},
			false,
		},
		{
			"example_3",
			args{[]int{5, 1, 2, 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompliance(tt.args.nums); got != tt.want {
				t.Errorf("isCompliance() = %v, want %v", got, tt.want)
			}
		})
	}
}
