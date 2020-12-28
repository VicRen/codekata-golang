package main

import (
	"reflect"
	"testing"
)

func Test_process(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{[]string{"HOW", "ARE", "YOU"}},
			[]string{"HAY", "ORO", "WEU"},
		},
		{
			"test2",
			args{[]string{"TO", "BE", "OR", "NOT", "TO", "BE"}},
			[]string{"TBONTB", "OEROOE", "   T"},
		},
		{
			"test3",
			args{[]string{"CONTEST", "IS", "COMING"}},
			[]string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transform(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test1",
			args{"HOW ARE YOU"},
			[]string{"HAY", "ORO", "WEU"},
		},
		{
			"test2",
			args{"TO BE OR NOT TO BE"},
			[]string{"TBONTB", "OEROOE", "   T"},
		},
		{
			"test3",
			args{"CONTEST IS COMING"},
			[]string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transform() = %v, want %v", got, tt.want)
			}
		})
	}
}
