package main

import (
	"reflect"
	"testing"
)

func Test_distinctPrimesOf(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"4",
			args{4},
			[]int{2},
		},
		{
			"14",
			args{14},
			[]int{2, 7},
		},
		{
			"644",
			args{644},
			[]int{2, 7, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctPrimesOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distinctPrimesOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
