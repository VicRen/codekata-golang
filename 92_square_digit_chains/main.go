package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0
	for i := 1; i < 10000000; i++ {
		if goChain(i) {
			count++
		}
	}
	fmt.Println(count)
}

func goChain(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, c := range s {
		i, _ := strconv.Atoi(string(c))
		sum += i * i
	}
	if sum == 89 {
		return true
	} else if sum == 1 {
		return false
	}
	return goChain(sum)
}
