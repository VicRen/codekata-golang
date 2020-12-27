package main

import "testing"

func Test_goChain(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{44},
			false,
		},
		{
			"test2",
			args{85},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goChain(tt.args.n); got != tt.want {
				t.Errorf("goChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
