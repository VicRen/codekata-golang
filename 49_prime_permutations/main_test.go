package main

import "testing"

func Test_isPerm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_1",
			args{123, 321},
			true,
		},
		{
			"test_2",
			args{123, 213},
			true,
		},
		{
			"test_3",
			args{123, 211},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isPerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
