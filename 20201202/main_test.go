package main

import "testing"

func TestRomaToDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{"III"},
			3,
		},
		{
			"test2",
			args{"IV"},
			4,
		},
		{
			"test3",
			args{"VI"},
			6,
		},
		{
			"test4",
			args{"LVIII"},
			58,
		},
		{
			"test5",
			args{"MCMXCIV"},
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomaToDigit(tt.args.s); got != tt.want {
				t.Errorf("RomaToDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
