package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiple(9, 9))
}

func multiple(n, m uint64) uint64 {
	if m == 1 {
		return n
	}
	return n + multiple(n, m-1)
}

func multiple2(n, m uint64) uint64 {
	sum := n
	m--
	if m > 0 {
		sum += multiple2(n, m)
	}
	return sum
}
