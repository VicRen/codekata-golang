package main

import (
	"fmt"
	"math"
)

func main() {
	primeList := generatePrimeList(10000)
	result := 1
	notFound := true
	for notFound {
		result += 2
		j := 0
		notFound = false
		for result >= primeList[j] {
			if isTwiceSquare(int64(result - primeList[j])) {
				notFound = true
				break
			}
			j++
		}
	}
	fmt.Println("result:", result)
}

func generatePrimeList(n int) []int {
	var ret []int
	for i := 1; i < n; i++ {
		if isPrime(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func isTwiceSquare(number int64) bool {
	x := math.Sqrt(float64(number / 2))
	return x == float64(int64(x))
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
