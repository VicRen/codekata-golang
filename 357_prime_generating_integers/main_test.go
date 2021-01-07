package main

import "testing"

func Test_isCompatible(t *testing.T) {
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
			args{30},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompatible(tt.args.n); got != tt.want {
				t.Errorf("isCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}
