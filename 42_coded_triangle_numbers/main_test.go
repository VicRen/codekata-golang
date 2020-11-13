package main

import (
	"reflect"
	"testing"
)

func Test_readWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"empty_input",
			args{""},
			nil,
		},
		{
			"one_name",
			args{"\"VIC\""},
			[]string{"VIC"},
		},
		{
			"two_names",
			args{"\"VIC\",\"COCO\""},
			[]string{"VIC", "COCO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readWords(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worthOf(t *testing.T) {
	tt := []struct {
		name string
		word string
		want int
	}{
		{
			"COLIN",
			"COLIN",
			53,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := worthOf(tc.word)
			if got != tc.want {
				t.Errorf("worthOf() = %d, want %d", got, tc.want)
			}
		})
	}
}

func Test_isTriangleNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{1},
			true,
		},
		{
			"28",
			args{28},
			true,
		},
		{
			"29",
			args{29},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTriangleNumber(tt.args.n); got != tt.want {
				t.Errorf("isTriangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
