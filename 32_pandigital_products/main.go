package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sum := 0
	for i := 0; i < 10000; i++ {
		if hasPandigitalProduct(i) {
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}

func hasPandigitalProduct(n int) bool {
	for i := 1; i <= n; i++ {
		for n%i == 0 && isPandigital(strconv.Itoa(n)+strconv.Itoa(i)+strconv.Itoa(n/i)) {
			fmt.Printf("Pandigital:%d = %d * %d\n", n, i, n/i)
			return true
		}
	}
	return false
}

func isPandigital(s string) bool {
	var ss []string
	for _, c := range s {
		ss = append(ss, string(c))
	}
	sort.Strings(ss)
	s = ""
	for _, c := range ss {
		s += c
	}
	return s == "123456789"
}
