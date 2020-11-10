package main

import (
	"fmt"
	"math"
)

func main() {
	max := 0
	fp := 0
	for p := 0; p <= 1000; p++ {
		n := len(findSolutionsFor(p))
		if n > max {
			max = n
			fp = p
		}
	}
	fmt.Println("max:", max, "fp", fp)
}

func findSolutionsFor(p int) [][]int {
	// a+b+c=p, a^2+b^2=c^2, a+b>c, c>a, c>b
	var ret [][]int
	for a := 0; a < p/2; a++ {
		for b := 0; b < a; b++ {
			cf := math.Sqrt(float64(a*a + b*b))
			c := int(cf)
			if cf-float64(c) == 0 && a+b+c == p {
				ret = append(ret, []int{b, a, c})
			}
		}
	}
	return ret
}
