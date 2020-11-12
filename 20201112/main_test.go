package main

import (
	"reflect"
	"testing"
)

func Test_sortFor(t *testing.T) {
	type args struct {
		n   int
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example_1",
			args{1, []int{5, 1, 6, 1}},
			[]int{1, 1, 5, 6},
		},
		{
			"example_2",
			args{2, []int{-1, 2, 3, 5, 2, 2}},
			[]int{2, 2, 2, -1, 3, 5},
		},
		{
			"example_3",
			args{1, []int{2, 3, 4, 6}},
			[]int{2, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortFor(tt.args.n, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortFor2(t *testing.T) {
	type args struct {
		n   int
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example_1",
			args{1, []int{5, 1, 6, 1}},
			[]int{1, 1, 5, 6},
		},
		{
			"example_2",
			args{2, []int{-1, 2, 3, 5, 2, 2}},
			[]int{2, 2, 2, -1, 3, 5},
		},
		{
			"example_3",
			args{1, []int{2, 3, 4, 6}},
			[]int{2, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortFor2(tt.args.n, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortFor2() = %v, want %v", got, tt.want)
			}
		})
	}
}
