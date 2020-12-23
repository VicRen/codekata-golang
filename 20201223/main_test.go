package main

import (
	"reflect"
	"testing"
)

func Test_find(t *testing.T) {
	type args struct {
		target     int
		candidates []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test1",
			args{
				8,
				[]int{10, 1, 2, 7, 6, 1, 5},
			},
			[][]int{
				{1, 7},
				{1, 2, 5},
				{2, 6},
				{1, 1, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.args.target, tt.args.candidates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
