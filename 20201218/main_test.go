package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{src: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			16,
		},
		{
			"test2",
			args{src: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			4,
		},
		{
			"test3",
			args{src: []string{"a", "aa", "aaa", "aaaa"}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.src); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasSameChar(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				s1: "abc",
				s2: "ade",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSameChar(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("hasSameChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
