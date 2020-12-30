package main

import (
	"reflect"
	"testing"
)

func Test_process(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				nums: []int{1, 2, 1},
			},
			[]int{2, -1, 2},
		},
		{
			"test2",
			args{
				nums: []int{1},
			},
			[]int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNextLarge(t *testing.T) {
	type args struct {
		index int
		nums  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				1,
				[]int{1, 2, 3},
			},
			3,
		},
		{
			"test2",
			args{
				0,
				[]int{1, 2, 3},
			},
			2,
		},
		{
			"test3",
			args{
				2,
				[]int{1, 2, 3},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNextLarge(tt.args.index, tt.args.nums)
			if got != tt.want {
				t.Errorf("findNextLarge() got = %v, want %v", got, tt.want)
			}
		})
	}
}
