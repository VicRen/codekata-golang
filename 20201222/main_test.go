package main

import "testing"

func Test_isCompatible(t *testing.T) {
	type args struct {
		a string
		b string
		c string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				"aabcc",
				"dbbca",
				"aadbbcbcac",
			},
			true,
		},
		{
			"test2",
			args{
				"aabcc",
				"dbbca",
				"aadbbbaccc",
			},
			false,
		},
		{
			"test3",
			args{
				"",
				"",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompatible(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("isCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}
