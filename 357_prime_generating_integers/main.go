package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 100000000; i++ {
		if isCompatible(i) {
			sum += i
		}
	}
	fmt.Println("sum=", sum)
}

func isCompatible(n int) bool {
	ds := listDivisors(n)
	for _, i := range ds {
		if !isPrime(i + n/i) {
			return false
		}
	}
	return true
}

func listDivisors(input int) []int {
	var ret []int
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			ret = append(ret, i)
		}
	}
	return ret
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
