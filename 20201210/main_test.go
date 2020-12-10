package main

import "testing"

func Test_replace(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				"G()(al)",
			},
			"Goal",
		},
		{
			"test2",
			args{
				"G()()()()(al)",
			},
			"Gooooal",
		},
		{
			"test3",
			args{
				"(al)G(al)()()G",
			},
			"alGalooG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replace(tt.args.src); got != tt.want {
				t.Errorf("replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
