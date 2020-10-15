package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumDigits(t *testing.T) {
	tt := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{
			"invalid",
			"abc",
			0,
			fmt.Errorf("invalid input: %s", "abc"),
		},
		{
			"invalid_2",
			"1a",
			0,
			fmt.Errorf("invalid input: %s", "1a"),
		},
		{
			"1",
			"1",
			1,
			nil,
		},
		{
			"123",
			"123",
			6,
			nil,
		},
		{
			"2345",
			"2345",
			14,
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SumDigits(tc.input)
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Errorf("SumDigits(%s) err=%v want err %v", tc.input, err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("SumDigits(%s)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
