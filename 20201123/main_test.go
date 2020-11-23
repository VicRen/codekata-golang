package main

import (
	"reflect"
	"testing"
)

//func Test_mergeSlice(t *testing.T) {
//	type args struct {
//		n []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{
//			"example_1",
//			args{[]int{1, 2, 0, 0}},
//			1200,
//		},
//		{
//			"example_2",
//			args{[]int{2, 7, 4}},
//			274,
//		},
//		{
//			"example_3",
//			args{[]int{2, 1, 5}},
//			215,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := mergeSlice(tt.args.n); got != tt.want {
//				t.Errorf("mergeSlice() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_stringToSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example_1",
			args{"1234"},
			[]int{1, 2, 3, 4},
		},
		{
			"example_2",
			args{"455"},
			[]int{4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution(t *testing.T) {
	type args struct {
		k int
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example_1",
			args{34, []int{1, 2, 0, 0}},
			[]int{1, 2, 3, 4},
		},
		{
			"example_2",
			args{181, []int{2, 7, 4}},
			[]int{4, 5, 5},
		},
		{
			"example_3",
			args{806, []int{2, 1, 5}},
			[]int{1, 0, 2, 1},
		},
		{
			"example_4",
			args{1, []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
			[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.k, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
