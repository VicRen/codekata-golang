package main

import (
	"reflect"
	"testing"
)

func Test_findSprialDiagonals(t *testing.T) {
	type args struct {
		width int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"1",
			args{1},
			[]int{1},
			false,
		},
		{
			"2",
			args{2},
			nil,
			true,
		},
		{
			"3",
			args{3},
			[]int{1, 3, 5, 7, 9},
			false,
		},
		{
			"5",
			args{5},
			[]int{1, 3, 5, 7, 9, 13, 17, 21, 25},
			false,
		},
		{
			"7",
			args{7},
			[]int{1, 3, 5, 7, 9, 13, 17, 21, 25, 31, 37, 43, 49},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSprialDiagonals(tt.args.width)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSprialDiagonals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSprialDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}
