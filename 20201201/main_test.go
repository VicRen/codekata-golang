package main

import "testing"

func Test_isEcho(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"aba"},
			true,
		},
		{
			"test2",
			args{"abca"},
			false,
		},
		{
			"test3",
			args{"abcba"},
			true,
		},
		{
			"test4",
			args{""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEcho(tt.args.s); got != tt.want {
				t.Errorf("isEcho() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlmostEcho(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"aba"},
			true,
		},
		{
			"test2",
			args{"abca"},
			true,
		},
		{
			"test3",
			args{"abcba"},
			true,
		},
		{
			"test4",
			args{""},
			false,
		},
		{
			"test5",
			args{"abcbba"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlmostEcho(tt.args.s); got != tt.want {
				t.Errorf("isAlmostEcho() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeIndex(t *testing.T) {
	type args struct {
		index int
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{0, "ab"},
			"b",
		},
		{
			"test2",
			args{0, ""},
			"",
		},
		{
			"test3",
			args{1, "abc"},
			"ac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeIndex(tt.args.index, tt.args.s); got != tt.want {
				t.Errorf("removeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
