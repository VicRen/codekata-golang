package main

import (
	"reflect"
	"testing"
)

func PrimeFactorsOf(n int) []int {
	ret := make([]int, 0)
	d := 2
	for d < n {
		for n%d == 0 {
			ret = append(ret, d)
			n /= d
		}
		d++
	}
	if n > 1 {
		ret = append(ret, n)
	}
	return ret
}

func TestPrimeFactorsOf(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  []int
	}{
		{"1", 1, []int{}},
		{"2", 2, []int{2}},
		{"3", 3, []int{3}},
		{"4", 4, []int{2, 2}},
		{"8", 8, []int{2, 2, 2}},
		{"9", 9, []int{3, 3}},
		{"a very large number", 2 * 3 * 17 * 71 * 73 * 101, []int{2, 3, 17, 71, 73, 101}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimeFactorsOf(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PrimeFactorsOf(%d) = %v, want %v\n", tc.input, got, tc.want)
			}
		})
	}
}
