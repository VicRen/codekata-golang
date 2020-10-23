package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_findPermutation(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tt := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			"1-2",
			args{1, 2},
			0,
			errors.New("m=2 cannot be larger than n=1"),
		},
		{
			"1-1",
			args{1, 1},
			1,
			nil,
		},
		{
			"3-3",
			args{3, 3},
			6,
			nil,
		},
		{
			"4-4",
			args{4, 4},
			24,
			nil,
		},
		{
			"3-2",
			args{3, 2},
			6,
			nil,
		},
		{
			"4-2",
			args{4, 2},
			12,
			nil,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := findPermutation(tc.args.n, tc.args.m)
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Errorf("findPermutation() = %v, want err %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("findPermutation() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_findNPermutation(t *testing.T) {
	type args struct {
		n    int
		src  []int
		facs []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"012-1",
			args{1, []int{0, 1, 2}, []int{1, 1, 2}},
			"012",
		},
		{
			"012-2",
			args{2, []int{0, 1, 2}, []int{1, 1, 2}},
			"021",
		},
		{
			"012-6",
			args{6, []int{0, 1, 2}, []int{1, 1, 2}},
			"210",
		},
		{
			"0123456-1",
			args{1, []int{0, 1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 6, 24, 120, 720}},
			"0123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNPermutation(tt.args.n, tt.args.src, tt.args.facs); got != tt.want {
				t.Errorf("findNPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
