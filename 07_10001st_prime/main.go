package main

import "fmt"

func main() {
	count := 0
	n := 0
	for {
		if count == 10001 {
			break
		}
		n++
		pf := PrimeFactorsOf(n)
		if len(pf) == 1 {
			count++
		}
	}
	fmt.Println("10001st prime is", n)
}

func PrimeFactorsOf(n int) []int {
	ret := make([]int, 0)
	divider := 2
	for divider < n {
		for n%divider == 0 {
			ret = append(ret, divider)
			n /= divider
		}
		divider++
	}
	if n > 1 {
		ret = append(ret, n)
	}
	return ret
}
