package main

import "testing"

func Test_findGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"4, 8",
			args{4, 8},
			4,
		},
		{
			"12, 48",
			args{12, 48},
			12,
		},
		{
			"12, 48",
			args{12, 48},
			12,
		},
		{
			"319, 377",
			args{319, 377},
			29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
