package main

func main() {

}

func isMountain(nums []int) bool {
	hasUp := false
	hasDown := false
	c := nums[0]
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if !hasUp && !hasDown && n > c {
			hasUp = true
		}
		if c > n {
			hasDown = true
		}
		if n > c && hasDown {
			return false
		}
		c = n
	}
	return hasDown && hasUp
}
