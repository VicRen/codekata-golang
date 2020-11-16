package main

import (
	"fmt"
	"math"
)

func main() {
	i := 144
	for {
		m := 2*i*i - i
		fmt.Println("i,m", i, m)
		if isPentagonalNumber(m) {
			fmt.Println("num:", m)
			break
		}
		i++
	}
}

func isPentagonalNumber(n int) bool {
	r := math.Sqrt(float64(1 + 24*n))
	if r-float64(int(r)) != 0 {
		return false
	}
	return int(r)%6 == 5
}

func isTriangleNumber(n int) bool {
	//tn = ½n(n+1)
	x := 2 * n
	res := x
	//牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	for i := res; i > 0; i-- {
		if i*(i+1) == x {
			return true
		}
	}
	return false
}
