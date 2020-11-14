package main

import "fmt"

var steps = []int{1, 2}

func main() {
	fmt.Println("count:", climbingStairs(100, steps))
}

func climbingStairs(n int, step []int) int {
	if n <= 1 {
		return 1
	}
	pre1 := 1
	pre2 := 1
	for i := 2; i <= n; i++ {
		cur := pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}
