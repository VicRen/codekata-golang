package main

import "testing"

func Test_maxCountWord(t *testing.T) {
	type args struct {
		s   string
		ban string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				s:   "Bob hit a ball, the hit BALL flew far after it was hit.",
				ban: "hit",
			},
			"ball",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCountWord(tt.args.s, tt.args.ban); got != tt.want {
				t.Errorf("maxCountWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCountWord2(t *testing.T) {
	type args struct {
		s   string
		ban string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				s:   "Bob hit a ball, the hit BALL flew far after it was hit.",
				ban: "hit",
			},
			"ball",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCountWord2(tt.args.s, tt.args.ban); got != tt.want {
				t.Errorf("maxCountWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
