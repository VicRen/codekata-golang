package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_process(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{
				n: 2,
			},
			[]string{"1/2"},
		},
		{
			"test2",
			args{
				n: 3,
			},
			[]string{"1/2", "1/3", "2/3"},
		},
		{
			"test3",
			args{
				n: 4,
			},
			[]string{"1/2", "1/3", "1/4", "2/3", "3/4"},
		},
		{
			"test4",
			args{
				n: 1,
			},
			nil,
		},
		{
			"test5",
			args{
				n: 9,
			},
			[]string{"1/2", "1/3", "1/4", "1/5", "1/6", "1/7", "1/8", "1/9", "2/3", "2/5", "2/7", "2/9", "3/4", "3/5", "3/7", "3/8", "4/5", "4/7", "4/9", "5/6", "5/7", "5/8", "5/9", "6/7", "7/8", "7/9", "8/9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := process(tt.args.n)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
