package main

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				"nlaebolko",
			},
			1,
		},
		{
			"test2",
			args{
				"loonbalxballpoon",
			},
			2,
		},
		{
			"test3",
			args{
				"leetcode",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}
