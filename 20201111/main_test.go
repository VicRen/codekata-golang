package main

import (
	"testing"
)

func Test_findSingleNumber(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"example_1",
			args{[]int{2, 2, 1}},
			1,
			false,
		},
		{
			"example_2",
			args{[]int{4, 1, 2, 1, 2}},
			4,
			false,
		},
		{
			"example_3",
			args{[]int{4, 4, 1, 2, 1, 2}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSingleNumber(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSingleNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSingleNumber2(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"example_1",
			args{[]int{2, 2, 1}},
			1,
			false,
		},
		{
			"example_2",
			args{[]int{4, 1, 2, 1, 2}},
			4,
			false,
		},
		{
			"example_3",
			args{[]int{4, 4, 1, 2, 1, 2}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSingleNumber2(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSingleNumber2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSingleNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSingleNumber3(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"example_1",
			args{[]int{2, 2, 1}},
			1,
			false,
		},
		{
			"example_2",
			args{[]int{4, 1, 2, 1, 2}},
			4,
			false,
		},
		{
			"example_3",
			args{[]int{4, 4, 1, 2, 1, 2}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSingleNumber3(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSingleNumber3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSingleNumber3() = %v, want %v", got, tt.want)
			}
		})
	}
}
