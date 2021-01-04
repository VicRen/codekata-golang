package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	solution()
}

func solution() {
	n := math.Sqrt(19293949596979899)
	t := math.Sqrt(10203040506070809)
	for i := n; i >= t; i-- {
		if i == 1389019170 {
			fmt.Println(i)
		}
		x := (&big.Int{}).SetInt64(int64(i))
		if isMatch(x.Mul(x, x)) {
			fmt.Println(x)
			fmt.Println("num is", int64(i)*10)
		}
	}
}

func isMatch(i *big.Int) bool {
	x := i.String()
	num := x[0]
	for i := 2; i < len(x); i += 2 {
		if x[i] == num+1 {
			num++
			continue
		}
		return false
	}
	return true
}
