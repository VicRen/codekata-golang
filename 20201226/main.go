package main

import (
	"math"
	"sort"
)

func main() {

}

func findThirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	if len(nums) < 3 {
		return nums[0]
	}
	v := nums[0]
	count := 0
	for _, n := range nums {
		v = n
		if count == 3 {
			break
		}
		if v > n {
			count++
		}
	}
	return v
}

func findThirdMax2(nums []int) int {
	vs := [3]int{math.MinInt64, math.MinInt64, math.MinInt64}
	found := false
	for _, n := range nums {
		if n == vs[0] || n == vs[1] || n == vs[2] {
			continue
		}
		if n > vs[0] {
			vs[0], vs[1], vs[2] = n, vs[0], vs[1]
		} else if n > vs[1] {
			vs[1], vs[2] = n, vs[1]
		} else if n > vs[2] {
			found = true
			vs[2] = n
		}
	}
	if found {
		return vs[2]
	}
	return vs[0]
}
