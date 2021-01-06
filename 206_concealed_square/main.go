package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	solution()
}

func solution() {
	ts := time.Now()
	n := math.Sqrt(19293949596979899)
	t := math.Sqrt(10203040506070809)
	for i := n; i >= t; i-- {
		x := (&big.Int{}).SetInt64(int64(i))
		if !isEndsWith37(x.String()) {
			continue
		}
		if isMatch(x.Mul(x, x)) {
			fmt.Println(x)
			fmt.Println("num is", int64(i)*10)
			fmt.Println(time.Since(ts))
			break
		}
	}
}

func isEndsWith37(s string) bool {
	if len(s) == 0 {
		return false
	}
	c := s[len(s)-1]
	if c == '3' || c == '7' {
		return true
	}
	return false
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
