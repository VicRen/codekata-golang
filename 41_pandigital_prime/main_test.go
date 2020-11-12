package main

import "testing"

func Test_isPandigital(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"912345678",
			args{"912345678"},
			true,
		},
		{
			"12345678",
			args{"12345678"},
			true,
		},
		{
			"812349756",
			args{"812349756"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPandigital(tt.args.s); got != tt.want {
				t.Errorf("isPandigital() = %v, want %v", got, tt.want)
			}
		})
	}
}
