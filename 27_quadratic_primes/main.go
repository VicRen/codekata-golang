package main

import "fmt"

func main() {
	maxN := 0
	maxA := 0
	maxB := 0

	cm := 0
	for a := -999; a < 1000; a += 2 {
		for b := 3; b < 1000; b += 2 {
			fmt.Println("a:", a, "b:", b)
			cm = findMaxN(a, b)
			if cm > maxN {
				maxN = cm
				maxA = a
				maxB = b
			}
		}
	}

	fmt.Println("product is:", maxA*maxB)
}

func findMaxN(a, b int) int {
	max := 0
	n := 2
	for isPrime(n*n + a*n + b) {
		if n > max {
			max = n
		}
		n++
	}
	return max
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	divider := 2
	for divider < n {
		if n%divider == 0 {
			return false
		}
		divider++
	}
	return true
}
