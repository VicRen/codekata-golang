package main

import (
	"reflect"
	"testing"
)

func TestFetchNextAdjacent(t *testing.T) {
	tt := []struct {
		name        string
		wantCount   int
		srcDigits   []int
		wantDigits  []int
		wantSrcLeft []int
	}{
		{
			"2 out of 3",
			2,
			[]int{1, 2, 3},
			[]int{2, 3},
			[]int{},
		},
		{
			"2 out of 2",
			2,
			[]int{1, 2},
			[]int{},
			[]int{1, 2},
		},
		{
			"5 out of 6",
			5,
			[]int{1, 2, 3, 4, 5, 6},
			[]int{2, 3, 4, 5, 6},
			[]int{},
		},
		{
			"2 out of 6",
			2,
			[]int{1, 2, 3, 4, 5, 6},
			[]int{2, 3},
			[]int{4, 5, 6},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, got2 := FetchNextAdjacent(tc.wantCount, tc.srcDigits)
			if !reflect.DeepEqual(got, tc.wantDigits) || !reflect.DeepEqual(got2, tc.wantSrcLeft) {
				t.Errorf("FetchNextAdjacent(%d, %v)=%v, %v, want %v, %v", tc.wantCount, tc.srcDigits, got, got2, tc.wantDigits, tc.wantSrcLeft)
			}
		})
	}
}
