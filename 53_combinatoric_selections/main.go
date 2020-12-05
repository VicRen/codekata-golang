package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	result := 0
	limit, _ := (&big.Int{}).SetString("1000000", 0)
	nLimit := 100
	for i := 1; i <= nLimit; i++ {
		for j := 0; j <= i; j++ {
			bi, _ := (&big.Int{}).SetString(strconv.Itoa(i), 0)
			fi := factorialOf(bi)
			bj, _ := (&big.Int{}).SetString(strconv.Itoa(j), 0)
			fj := factorialOf(bj)
			bk, _ := (&big.Int{}).SetString(strconv.Itoa(i-j), 0)
			fk := factorialOf(bk)
			mik := fj.Mul(fj, fk)
			dvk := mik.Div(fi, mik)
			if x := dvk.Add(dvk, limit.Neg(limit)); x.Int64() >= 0 {
				result++
			}
		}
	}
	fmt.Println(result)
}

func factorialOf(n *big.Int) *big.Int {
	one, _ := (&big.Int{}).SetString("1", 0)
	if n.Int64() <= 2 {
		return one
	}
	p, _ := (&big.Int{}).SetString("1", 0)
	i, _ := (&big.Int{}).SetString("2", 0)
	for {
		p = p.Mul(p, i)
		if i.String() == n.String() {
			break
		}
		i = i.Add(i, one)
	}
	return p
}
