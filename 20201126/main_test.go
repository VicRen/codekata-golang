package main

import "testing"

func Test_findMaxSubLen(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example_1",
			args{"abccefg"},
			4,
		},
		{
			"example_2",
			args{"bbbbb"},
			1,
		},
		{
			"example_3",
			args{"pwwkew"},
			3,
		},
		{
			"example_4",
			args{""},
			0,
		},
		{
			"example_5",
			args{"2"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSubLen(tt.args.src); got != tt.want {
				t.Errorf("findMaxSubLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxSubLen2(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example_1",
			args{"abccefg"},
			4,
		},
		{
			"example_2",
			args{"bbbbb"},
			1,
		},
		{
			"example_3",
			args{"pwwkew"},
			3,
		},
		{
			"example_4",
			args{""},
			0,
		},
		{
			"example_5",
			args{"2"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSubLen2(tt.args.src); got != tt.want {
				t.Errorf("findMaxSubLen2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasValue(t *testing.T) {
	type args struct {
		n   uint8
		src []uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{1, []uint8{1, 2}},
			0,
		},
		{
			"test2",
			args{1, []uint8{3, 2}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOfValue(tt.args.n, tt.args.src); got != tt.want {
				t.Errorf("indexOfValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
