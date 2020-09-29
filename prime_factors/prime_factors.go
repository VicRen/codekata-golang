package main

import "fmt"

func main() {
	fmt.Println("prime number up to 100 is:")
	for i := 0; i < 101; i++ {
		pf := PrimeFactorsOf(i)
		if len(pf) < 2 {
			fmt.Println(i)
		}
	}
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
