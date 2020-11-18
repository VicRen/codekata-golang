package main

import (
	"reflect"
	"testing"
)

func Test_isTwiceSquare(t *testing.T) {
	type args struct {
		number int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"8",
			args{8},
			true,
		},
		{
			"10",
			args{10},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTwiceSquare(tt.args.number); got != tt.want {
				t.Errorf("isTwiceSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePrimeList(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"10",
			args{10},
			[]int{2, 3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePrimeList(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePrimeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
