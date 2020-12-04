package main

import (
	"sort"
)

func main() {

}

func countMove(nums []int) int {
	count := 0
	l := len(nums)
	for {
		if isAllSame(nums) {
			return count
		}
		sort.Ints(nums)
		for i := 0; i < l-1; i++ {
			nums[i] += 1
		}
		count++
	}
}

func isAllSame(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	base := nums[0]
	for _, n := range nums {
		if base != n {
			return false
		}
	}
	return true
}
