package main

import "testing"

func Test_freqAlphabet(t *testing.T) {
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
			args{
				"10#11#12",
			},
			"jkab",
		},
		{
			"test2",
			args{
				"1326#",
			},
			"acz",
		},
		{
			"test3",
			args{
				"25#",
			},
			"y",
		},
		{
			"test4",
			args{
				"12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			},
			"abcdefghijklmnopqrstuvwxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabet(tt.args.s); got != tt.want {
				t.Errorf("freqAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toChar(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			"test1",
			args{
				"1",
			},
			'a',
		},
		{
			"test2",
			args{
				"2",
			},
			'b',
		},
		{
			"test3",
			args{
				"9",
			},
			'i',
		},
		{
			"test4",
			args{
				"10",
			},
			'j',
		},
		{
			"test5",
			args{
				"10",
			},
			'j',
		},
		{
			"test6",
			args{
				"26",
			},
			'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toChar(tt.args.c); got != tt.want {
				t.Errorf("toChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
