package main

import (
	"math"
)

func main() {
}

func isPrime(n int) bool {
	res := int(math.Sqrt(float64(n)))
	divider := 2
	for divider <= res {
		if n%divider == 0 {
			return false
		}
		divider++
	}
	return true
}
