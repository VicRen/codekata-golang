package main

import "testing"

func Test_maskPII(t *testing.T) {
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
			args{"LeetCode@LeetCode.com"},
			"l*****e@leetcode.com",
		},
		{
			"test2",
			args{"AB@qq.com"},
			"a*****b@qq.com",
		},
		{
			"test3",
			args{"1(234)567-890"},
			"***-***-7890",
		},
		{
			"test4",
			args{"86-(10)12345678"},
			"+**-***-***-5678",
		},
		{
			"test5",
			args{"+(501321)-50-23431"},
			"+***-***-***-3431",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskPII(tt.args.s); got != tt.want {
				t.Errorf("maskPII() = %v, want %v", got, tt.want)
			}
		})
	}
}
