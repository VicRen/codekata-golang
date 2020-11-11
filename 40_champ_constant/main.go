package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	p := 1
	for i := 0; i < 7; i++ {
		n := int(math.Pow(10, float64(i)))
		fmt.Println(n)
		p *= getDigit(n)
	}
	fmt.Println("num=", p)
}

func getDigit(n int) int {
	var r, s, k int
	for s < n {
		r = s
		k++
		s += k * 9 * int(math.Pow(10, float64(k-1)))
	}
	h := n - r - 1
	t := int(math.Pow(10, float64(k-1))) + h/k
	p := h % k
	for i, s := range strconv.Itoa(t) {
		if i == p {
			ret, _ := strconv.Atoi(string(s))
			return ret
		}
	}
	return 0
}
