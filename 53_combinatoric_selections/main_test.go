package main

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_factorialOf(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			"test1",
			args{(&big.Int{}).SetInt64(0)},
			(&big.Int{}).SetInt64(1),
		},
		{
			"test2",
			args{(&big.Int{}).SetInt64(3)},
			(&big.Int{}).SetInt64(6),
		},
		{
			"test2",
			args{(&big.Int{}).SetInt64(5)},
			(&big.Int{}).SetInt64(120),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factorialOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
