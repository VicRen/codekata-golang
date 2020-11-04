package main

import (
	"reflect"
	"testing"
)

func Test_factorialOf(t *testing.T) {
	type args struct {
		n int
	}
	tt := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{1},
			1,
		},
		{
			"2",
			args{2},
			2,
		},
		{
			"4",
			args{4},
			24,
		},
		{
			"10",
			args{10},
			3628800,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := factorialOf(tc.args.n); got != tc.want {
				t.Errorf("factorialOf() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_splitNumber(t *testing.T) {
	type args struct {
		n int
	}
	tt := []struct {
		name string
		args args
		want []int
	}{
		{
			"13",
			args{13},
			[]int{1, 3},
		},
		{
			"123",
			args{123},
			[]int{1, 2, 3},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := splitNumber(tc.args.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("splitNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_sumNumberFactorials(t *testing.T) {
	type args struct {
		ns []int
	}
	tt := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{1}},
			1,
		},
		{
			"12",
			args{[]int{1, 2}},
			3,
		},
		{
			"123",
			args{[]int{1, 2, 3}},
			9,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := sumNumberFactorials(tc.args.ns); got != tc.want {
				t.Errorf("sumNumberFactorials() = %v, want %v", got, tc.want)
			}
		})
	}
}
