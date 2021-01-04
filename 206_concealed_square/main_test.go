package main

import (
	"math/big"
	"testing"
)

func Test_isMatch(t *testing.T) {
	fn, _ := (&big.Int{}).SetString("10203040506070809", 0)
	type args struct {
		i *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				fn,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.i); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
