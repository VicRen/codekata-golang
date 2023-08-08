package main

import "testing"

func binarySearch(arr []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, l, mid-1, target)
	} else {
		return binarySearch(arr, mid+1, r, target)
	}
}

func TestBinarySearch(t *testing.T) {
	type input struct {
		arr    []int
		l      int
		r      int
		target int
	}
	tt := []struct {
		name  string
		input input
		want  int
	}{
		{"first", input{
			arr:    []int{2, 3, 4, 5, 7},
			l:      0,
			r:      4,
			target: 5,
		}, 3},
		{"second", input{
			arr:    []int{2, 3, 4, 5, 7},
			l:      0,
			r:      4,
			target: 2,
		}, 0},
		{"third", input{
			arr:    []int{2, 3, 4, 5, 7},
			l:      0,
			r:      4,
			target: 7,
		}, 4},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := binarySearch(tc.input.arr, tc.input.l, tc.input.r, tc.input.target)
			if got != tc.want {
				t.Errorf("binarySearch(%v) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

func TestSum(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{"1", 1, 1},
		{"2", 2, 3},
		{"3", 3, 6},
		{"4", 4, 10},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := sum(tc.input)
			if got != tc.want {
				t.Errorf("sum(%d)=%d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
