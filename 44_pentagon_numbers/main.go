package main

import (
	"fmt"
	"math"
)

func main() {
	j := 1
	f := true
	for f {
		a := j * (3*j - 1) / 2
		for k := j + 1; k < 5000; k++ {
			b := k * (3*k - 1) / 2
			d := b - a
			if isPentagonalNumber(d) && isPentagonalNumber(a+b) {
				fmt.Println(j, k, a, b, d, a+b)
				f = false
				break
			}
		}
		j++
	}
}

func pentagonalNumberN(n int) int {
	return n * (3*n - 1) / 2
}

func isPentagonalNumber(n int) bool {
	r := math.Sqrt(float64(1 + 24*n))
	if r-float64(int(r)) != 0 {
		return false
	}
	return int(r)%6 == 5
}
