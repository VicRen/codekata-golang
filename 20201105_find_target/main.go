package main

import "fmt"

var nums = []int{1, 2, 7, 9, 13, 57, 36, 26, 55, 11, 6, 3, 89, 36, 75, 36, 76, 95, 98, 101, 320, 520, 85, 36, 62, 49, 96, 1}

func main() {
	fmt.Println("num=", findTargetInNums(35, nums))
	fmt.Println(findTargetInNums2(35, nums))
}

func findTargetInNums(target int, nums []int) []int {
	for x, i := range nums {
		if i >= target {
			continue
		}
		for y, j := range nums {
			if i == j {
				continue
			}
			if i+j == target {
				return []int{x, y}
			}
		}
	}
	return nil
}

func findTargetInNums2(target int, nums []int) []ret {
	var r []ret
	for x, i := range nums {
		if i >= target {
			continue
		}
		for y, j := range nums {
			if i == j {
				continue
			}
			if i+j == target {
				r = append(r, ret{i, x})
				r = append(r, ret{j, y})
				return r
			}
		}
	}
	return r
}

type ret struct {
	num       int
	subscript int
}

func (r ret) String() string {
	return fmt.Sprintf("num=%d-sub=%d", r.num, r.subscript)
}
