package main

import (
	"testing"
)

func Test_calWeek(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tt := []struct {
		name string
		args args
		want int
	}{
		{
			"1900-01-01",
			args{
				1900,
				01,
				01,
			},
			1,
		},
		{
			"2020-10-19",
			args{
				2020,
				10,
				19,
			},
			1,
		},
		{
			"2020-10-21",
			args{
				2020,
				10,
				21,
			},
			3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := calWeek(tc.args.year, tc.args.month, tc.args.day)
			if got != tc.want {
				t.Errorf("calWeek = %d, want %d", got, tc.want)
			}
		})
	}
}

func Test_isLeapYear(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  bool
	}{
		{
			"1900",
			1900,
			false,
		},
		{
			"2000",
			2000,
			true,
		},
		{
			"2020",
			2020,
			true,
		},
		{
			"2021",
			2021,
			false,
		},
		{
			"1800",
			1800,
			false,
		},
		{
			"1600",
			1600,
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := isLeapYear(tc.input)
			if got != tc.want {
				t.Errorf("isLeapYear(%d) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
