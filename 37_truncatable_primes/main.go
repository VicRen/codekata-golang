package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	start := 8
	count := 0
	for {
		if count == 11 {
			break
		}
		if isTruncatablePrimesL2R(start) && isTruncatablePrimesR2L(start) {
			count++
			fmt.Println("tp:", start)
			sum += start
		}
		start++
	}
	fmt.Println("sum:", sum)
}

func isTruncatablePrimesL2R(n int) bool {
	s := strconv.Itoa(n)
	for len(s) > 0 {
		x, _ := strconv.Atoi(s)
		if !isPrime(x) {
			return false
		}
		s = s[1:]
	}
	return true
}

func isTruncatablePrimesR2L(n int) bool {
	s := strconv.Itoa(n)
	for len(s) > 0 {
		l := len(s)
		x, _ := strconv.Atoi(s)
		if !isPrime(x) {
			return false
		}
		s = s[:l-1]
	}
	return true
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	res := n
	//牛顿法求平方根
	for res*res > n {
		res = (res + n/res) / 2
	}
	divider := 2
	for divider <= res {
		if n%divider == 0 {
			return false
		}
		divider++
	}
	return true
}
