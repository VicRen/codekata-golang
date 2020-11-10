package main

import (
	"reflect"
	"testing"
)

func Test_findSolutionsFor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"120",
			args{120},
			[][]int{{30, 40, 50}, {24, 45, 51}, {20, 48, 52}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSolutionsFor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSolutionsFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
