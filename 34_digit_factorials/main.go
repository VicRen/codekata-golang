package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	max := factorialOf(9) * 10
	for i := 3; i < max; i++ {
		if sumNumberFactorials(splitNumber(i)) == i {
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}

func sumNumberFactorials(ns []int) int {
	sum := 0
	for _, n := range ns {
		sum += factorialOf(n)
	}
	return sum
}

func splitNumber(n int) []int {
	var ret []int
	str := strconv.Itoa(n)
	for _, c := range str {
		i, _ := strconv.Atoi(string(c))
		ret = append(ret, i)
	}
	return ret
}

func factorialOf(n int) int {
	p := 1
	for i := n; i > 0; i-- {
		p *= i
	}
	return p
}
