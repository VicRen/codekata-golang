package main

import (
	"testing"
)

func Test_isTruncatablePrimesL2R(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"3797",
			args{3797},
			true,
		},
		{
			"797",
			args{797},
			true,
		},
		{
			"97",
			args{97},
			true,
		},
		{
			"96",
			args{96},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTruncatablePrimesL2R(tt.args.n); got != tt.want {
				t.Errorf("isTruncatablePrimesL2R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTruncatablePrimesR2L(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"3797",
			args{3797},
			true,
		},
		{
			"379",
			args{379},
			true,
		},
		{
			"37",
			args{37},
			true,
		},
		{
			"96",
			args{96},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTruncatablePrimesR2L(tt.args.n); got != tt.want {
				t.Errorf("isTruncatablePrimesR2L() = %v, want %v", got, tt.want)
			}
		})
	}
}
