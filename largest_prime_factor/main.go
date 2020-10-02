package main

import "fmt"

func main() {
	fmt.Println("the largest prime factor of 600851475143 is", LargestPrimeFactor(600851475143))
}

func LargestPrimeFactor(n int) int {
	divider := 2
	for divider < n {
		for n%divider == 0 {
			n /= divider
		}
		divider++
	}
	if n > 1 {
		return n
	} else {
		divider--
		return divider
	}
}
