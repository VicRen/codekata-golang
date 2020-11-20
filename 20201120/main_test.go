package main

import (
	"reflect"
	"testing"
)

func Test_turnOver(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"example_1",
			args{"ab-cd"},
			"dc-ba",
		},
		{
			"example_2",
			args{"a-bC-dEf-ghIj"},
			"j-Ih-gfE-dCba",
		},
		{
			"example_3",
			args{"Test1ng-Leet=code-Q!"},
			"Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := turnOver(tt.args.s); got != tt.want {
				t.Errorf("turnOver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLetter(t *testing.T) {
	type args struct {
		s int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"a",
			args{'a'},
			true,
		},
		{
			"U",
			args{'U'},
			true,
		},
		{
			"-",
			args{'-'},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLetter(tt.args.s); got != tt.want {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		a     []int32
		index int
		value int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			"test_1",
			args{
				[]int32{1, 2, 3},
				1,
				32,
			},
			[]int32{1, 32, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.a, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
