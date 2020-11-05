package main

import (
	"reflect"
	"testing"
)

func Test_findTargetInNums(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"3_in_1_2",
			args{3, []int{1, 2}},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetInNums(tt.args.target, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTargetInNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetInNums2(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want []ret
	}{
		{
			"3_in_1_2",
			args{3, []int{1, 2}},
			[]ret{{1, 0}, {2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetInNums2(tt.args.target, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTargetInNums2() = %v, want %v", got, tt.want)
			}
		})
	}
}
