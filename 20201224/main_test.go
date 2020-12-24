package main

import "testing"

func Test_isAlphabetAndDigit(t *testing.T) {
	type args struct {
		c uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{c: 'a'},
			true,
		},
		{
			"test2",
			args{c: '1'},
			true,
		},
		{
			"test3",
			args{c: '-'},
			false,
		},
		{
			"test4",
			args{c: ','},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlphabetAndDigit(tt.args.c); got != tt.want {
				t.Errorf("isAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			args{s: "A man, a plan, a canal: Panama"},
			true,
		},
		{
			"test2",
			args{s: "race a car"},
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
