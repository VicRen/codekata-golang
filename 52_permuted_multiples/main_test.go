package main

import "testing"

func Test_isSameDigits(t *testing.T) {
	type args struct {
		value [6]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{[6]int{2, 4, 6, 8, 10, 12}},
			false,
		},
		{
			"test2",
			args{[6]int{142857, 285714, 428571, 571428, 714285, 857142}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameDigits(tt.args.value); got != tt.want {
				t.Errorf("isSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
