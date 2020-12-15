package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				"()",
			},
			true,
		},
		{
			"test2",
			args{
				"()[]{}",
			},
			true,
		},
		{
			"test3",
			args{
				"(]",
			},
			false,
		},
		{
			"test4",
			args{
				"([)]",
			},
			false,
		},
		{
			"test5",
			args{
				"{[]}",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
