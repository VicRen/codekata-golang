package main

import "testing"

func Test_sumCoins(t *testing.T) {
	type args struct {
		target int
	}
	tt := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{1},
			1,
		},
		{
			"2",
			args{2},
			2,
		},
		{
			"3",
			args{3},
			2,
		},
		{
			"4",
			args{4},
			3,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := sumCoins(tc.args.target); got != tc.want {
				t.Errorf("sumCoins() = %v, want %v", got, tc.want)
			}
		})
	}
}
