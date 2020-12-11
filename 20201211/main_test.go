package main

import "testing"

func Test_process(t *testing.T) {
	type args struct {
		k int
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
				2,
				"abcdefg",
			},
			"bacdfeg",
		},
		{
			"test2",
			args{
				4,
				"abcdefg",
			},
			"dcbaefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.k, tt.args.s); got != tt.want {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
