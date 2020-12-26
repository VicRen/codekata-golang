package main

import "testing"

func Test_findThirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"in:[3,2,1],out:1",
			args{
				[]int{3, 2, 1},
			},
			1,
		},
		{
			"in:[1,2],out:2",
			args{
				[]int{1, 2},
			},
			2,
		},
		{
			"in:[2,2,3,1],out:1",
			args{
				[]int{2, 2, 3, 1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThirdMax(tt.args.nums); got != tt.want {
				t.Errorf("findThirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findThirdMax2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"in:[3,2,1],out:1",
			args{
				[]int{3, 2, 1},
			},
			1,
		},
		{
			"in:[1,2],out:2",
			args{
				[]int{1, 2},
			},
			2,
		},
		{
			"in:[2,2,3,1],out:1",
			args{
				[]int{2, 2, 3, 1},
			},
			1,
		},
		{
			"in:[2,2,3],out:3",
			args{
				[]int{2, 2, 3},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThirdMax2(tt.args.nums); got != tt.want {
				t.Errorf("findThirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
