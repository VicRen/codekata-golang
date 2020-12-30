package main

func main() {

}

func process(nums []int) []int {
	var ret []int
	c := 0
	for {
		if len(ret) == len(nums) {
			break
		}
		n := findNextLarge(c, nums)
		ret = append(ret, n)
		c++
	}
	return ret
}

func findNextLarge(index int, nums []int) int {
	if index > len(nums)-1 {
		return -1
	}
	src := append(nums[index+1:], nums[:index+1]...)
	for i, c := range src {
		if c > nums[index] {
			x := index + i + 1
			if x > len(nums) {
				x -= len(nums)
			}
			return c
		}
	}
	return -1
}
