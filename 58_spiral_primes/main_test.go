package main

import "testing"

func Test_primeRatio(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test1",
			args{[]int{2, 3, 5, 8}},
			0.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primeRatio(tt.args.nums); got != tt.want {
				t.Errorf("primeRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
