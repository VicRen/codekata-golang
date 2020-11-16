package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 7654321; i >= 2143; i -= 2 {
		if !isPandigital(strconv.Itoa(i)) {
			continue
		}
		fmt.Println(i)
		if isPrime(i) {
			fmt.Println("largest pandigital prime:", i)
			break
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	res := n
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

func isPandigital(s string) bool {
	l := len(s)
	for i := 1; i <= l; i++ {
		if !strings.Contains(s, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}
