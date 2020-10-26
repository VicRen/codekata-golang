package main

import (
	"reflect"
	"testing"
)

func Test_findDecimalOfOneNth(t *testing.T) {
	type args struct {
		n int
	}
	tt := []struct {
		name string
		args args
		want int
	}{
		{
			"2",
			args{2},
			1,
		},
		{
			"4",
			args{4},
			2,
		},
		{
			"6",
			args{6},
			1,
		},
		{
			"7",
			args{7},
			6,
		},
		{
			"10",
			args{10},
			1,
		},
		{
			"100",
			args{100},
			2,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := findCycleLengthOfOneNth(tc.args.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("findCycleLengthOfOneNth() = %v, want %v", got, tc.want)
			}
		})
	}
}
