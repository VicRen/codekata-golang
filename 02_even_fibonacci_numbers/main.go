package main

import (
	"fmt"
)

func main() {
	end := 4000000
	nums := GetFibonacciTill(end)
	fmt.Println("the fibonacci below 4000000 is", nums)
	sum := 0
	for _, n := range nums {
		if IsEvenValue(n) {
			sum += n
		}
	}
	fmt.Println("event-value sum of numbers in 4000000 is", sum)
}

func GetFibonacciTill(n int) []int {
	var ret []int
	next := 1
	last := 1
	for next <= n {
		add := next
		ret = append(ret, add)
		next += last
		last = add
	}
	return ret
}

func IsEvenValue(n int) bool {
	return n%2 == 0
}
