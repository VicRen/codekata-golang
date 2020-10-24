package main

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		index int
	}
	tt := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			"1",
			args{1},
			(&big.Int{}).SetBits([]big.Word{1}),
		},
		{
			"10",
			args{10},
			(&big.Int{}).SetBits([]big.Word{big.Word(55)}),
		},
		{
			"12",
			args{12},
			(&big.Int{}).SetBits([]big.Word{big.Word(144)}),
		},
		{
			"13",
			args{13},
			(&big.Int{}).SetBits([]big.Word{big.Word(233)}),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := fibonacci(tc.args.index); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tc.want)
			}
		})
	}
}
