package main

import "sort"

func main() {

}

func find(target int, candidates []int) [][]int {
	l := len(candidates)
	var ret [][]int
	sort.Ints(candidates)
	for k, n := range candidates {
		for j := 1; j < l-k-1; j++ {
			sum := n
			nums := []int{n}
			for i := k + j; i < l; i++ {
				num := candidates[i]
				sum += num
				if sum > target {
					break
				}
				nums = append(nums, num)
				if sum == target {
					ret = append(ret, nums)
					break
				}
			}
		}
	}
	return ret
}
