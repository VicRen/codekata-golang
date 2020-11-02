package main

import (
	"fmt"
	"strconv"
)

var nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	var facs []int
	for _, i := range nums {
		n := findFacs(i)
		facs = append(facs, n)
	}
	fmt.Println("facs:", facs)

	fmt.Println("the millionth is:", findNPermutation(1e6, nums, facs))
}

func findNPermutation(n int, src []int, facs []int) string {
	n--
	var ret string
	for i := len(src) - 1; i >= 0; i-- {
		t := n / facs[i]
		n %= facs[i]
		ret += strconv.Itoa(src[t])
		src = append(src[:t], src[t+1:]...)
	}
	return ret
}

func findFacs(n int) int {
	p := 1
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}
