package main

import "testing"

func Test_say(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				num: "",
			},
			"1",
		},
		{
			"test2",
			args{
				num: "1",
			},
			"11",
		},
		{
			"test3",
			args{
				num: "3322251",
			},
			"23321511",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := say(tt.args.num); got != tt.want {
				t.Errorf("say() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_process(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				1,
			},
			"1",
		},
		{
			"test2",
			args{
				2,
			},
			"11",
		},
		{
			"test3",
			args{
				3,
			},
			"21",
		},
		{
			"test4",
			args{
				4,
			},
			"1211",
		},
		{
			"test4",
			args{
				5,
			},
			"111221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.n); got != tt.want {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
