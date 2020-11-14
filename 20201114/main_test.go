package main

import "testing"

func Test_climbingStairs(t *testing.T) {
	type args struct {
		n    int
		step []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{2, []int{1, 2}},
			2,
		},
		{
			"3",
			args{3, []int{1, 2}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingStairs(tt.args.n, tt.args.step); got != tt.want {
				t.Errorf("climbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
