package main

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tt := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{1},
			false,
		},
		{
			"2",
			args{2},
			true,
		},
		{
			"4",
			args{4},
			false,
		},
		{
			"17",
			args{17},
			true,
		},
		{
			"37",
			args{37},
			true,
		},
		{
			"73",
			args{73},
			true,
		},
		{
			"13 * 3 * 5 * 5 * 17 * 31 * 737",
			args{3 * 3 * 5 * 5 * 17 * 31 * 73},
			false,
		},
		{
			"large_one",
			args{2876*2876 - 63*2876 + 977},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := isPrime(tc.args.n); got != tc.want {
				t.Errorf("isPrime() = %v, want %v", got, tc.want)
			}
		})
	}
}
