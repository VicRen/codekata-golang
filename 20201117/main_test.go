package main

import (
	"reflect"
	"testing"
)

func Test_mergeSlices(t *testing.T) {
	type args struct {
		src [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case_1",
			args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			[]int{1, 2, 3, 6, 8, 10, 15, 18},
		},
		{
			"case_2",
			args{[][]int{{1, 4}, {4, 5}}},
			[]int{1, 4, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAndSortSlices(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeAndSortSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pickInterval(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case_1",
			args{[]int{1, 2, 3, 6, 8, 10, 15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"case_2",
			args{[]int{1, 1}},
			[][]int{{1, 1}},
		},
		{
			"case_3",
			args{[]int{1, 2, 2, 3, 6, 8, 10, 15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"case_4",
			args{[]int{1, 4, 4, 5}},
			[][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickInterval(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pickInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		src [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case_1",
			args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"case_2",
			args{[][]int{{1, 4}, {4, 5}}},
			[][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
