package main

import (
	"testing"
)

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
			false,
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

func Test_hasPandigitalProduct(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"7254",
			args{7254},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPandigitalProduct(tt.args.n); got != tt.want {
				t.Errorf("hasPandigitalProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
