package main

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				[]int{1, 2, 3, 4},
			},
			[]int{2, 4, 4, 4},
		},
		{
			"test2",
			args{
				[]int{1, 1, 2, 3},
			},
			[]int{1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
