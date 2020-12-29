package main

import "testing"

func Test_process(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"ab#c"},
			"ac",
		},
		{
			"test2",
			args{"ad#c"},
			"ac",
		},
		{
			"test3",
			args{"ab##"},
			"",
		},
		{
			"test4",
			args{"c#d#"},
			"",
		},
		{
			"test5",
			args{"a##c"},
			"c",
		},
		{
			"test6",
			args{"#a#c"},
			"c",
		},
		{
			"test7",
			args{"a#c"},
			"c",
		},
		{
			"test8",
			args{"b"},
			"b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.s); got != tt.want {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEqual(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				"ab#c",
				"ad#c",
			},
			true,
		},
		{
			"test2",
			args{
				"ab##",
				"c#d#",
			},
			true,
		},
		{
			"test3",
			args{
				"a##c",
				"#a#c",
			},
			true,
		},
		{
			"test3",
			args{
				"a#c",
				"b",
			},
			false,
		},
		{
			"test4",
			args{
				"ac##",
				"a##",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
