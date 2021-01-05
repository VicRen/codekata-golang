package main

import "testing"

func Test_countBalance(t *testing.T) {
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
			args{
				"RLRRLLRLRL",
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalance(tt.args.s); got != tt.want {
				t.Errorf("countBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
