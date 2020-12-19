package main

import (
	"fmt"
	"math"
)

func main() {
	c := 0.
	for n := 1.; n <= 9.; n++ {
		c += math.Floor(math.Ln10 / math.Log(10/n))
	}
	fmt.Println("solution=", c)
}

func nPower(n, num int) int64 {
	p := int64(1)
	for i := 0; i < n; i++ {
		p *= int64(num)
	}
	return p
}
