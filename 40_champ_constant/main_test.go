package main

import "testing"

func Test_getDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"12",
			args{12},
			1,
		},
		{
			"13",
			args{13},
			1,
		},
		{
			"15",
			args{15},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDigit(tt.args.n); got != tt.want {
				t.Errorf("getDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
