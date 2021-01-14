package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{
				1,
			},
			[]string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readBinaryWatch(tt.args.num)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
