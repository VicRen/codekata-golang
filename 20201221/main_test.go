package main

import "testing"

func Test_isMountain(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{[]int{2, 1}},
			false,
		},
		{
			"test2",
			args{[]int{3, 5, 5}},
			false,
		},
		{
			"test3",
			args{[]int{0, 3, 2, 1}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMountain(tt.args.nums); got != tt.want {
				t.Errorf("isMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
