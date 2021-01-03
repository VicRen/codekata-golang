package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				"aaaabbbbcccc",
			},
			"abccbaabccba",
		},
		{
			"test2",
			args{
				"rat",
			},
			"art",
		},
		{
			"test3",
			args{
				"leetcode",
			},
			"cdelotee",
		},
		{
			"test4",
			args{
				"ggggggg",
			},
			"ggggggg",
		},
		{
			"test5",
			args{
				"spo",
			},
			"ops",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
