package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 11
	for {
		fb := fibonacci(n)
		fmt.Println("index =", n, "len =", len(fb.String()))
		if numberDigits(fb) >= 1000 {
			break
		}
		n++
	}

	fmt.Println("n is", n)
}

func fibonacci(index int) *big.Int {
	if index < 3 {
		ret := (&big.Int{}).SetUint64(1)
		return ret
	}
	fn := &big.Int{}
	fn1 := (&big.Int{}).SetInt64(1)
	fn2 := (&big.Int{}).SetInt64(1)
	for i := 3; i <= index; i++ {
		fn.Set(fn.Add(fn1, fn2))
		fn2.Set(fn1)
		fn1.Set(fn)
	}
	return fn
}

func numberDigits(num *big.Int) int {
	return len(num.String())
}
