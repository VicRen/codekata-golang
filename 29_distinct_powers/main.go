package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	m := map[string]struct{}{}
	count := 0
	for i := 2; i <= 100; i++ {
		for j := 2; j <= 100; j++ {
			c := 1
			p, _ := (&big.Int{}).SetString(strconv.Itoa(1), 0)
			t, _ := (&big.Int{}).SetString(strconv.Itoa(i), 0)
			for c <= j {
				p = p.Mul(p, t)
				c++
			}
			fmt.Println(i, j, p.String())
			count++
			if _, ok := m[p.String()]; !ok {
				m[p.String()] = struct{}{}
			}
		}
	}
	fmt.Println("len:", len(m), "count:", count)
}
