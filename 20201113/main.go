package main

import "fmt"

func main() {
	fmt.Println(isCompliance([]int{2, 3, 3, 2, 4}))
	fmt.Println(isCompliance([]int{3, 3, 2, 4}))
	fmt.Println(isCompliance([]int{4, 5, 1, 2}))
}

func isCompliance(nums []int) bool {
	sc := 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				if sc > 1 {
					return false
				}
				sc++
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return true
}
