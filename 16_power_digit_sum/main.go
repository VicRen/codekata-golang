package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	n := &big.Int{}
	n.SetInt64(1024)
	p := (&big.Int{}).SetInt64(1)
	for i := 0; i < 100; i++ {
		p = p.Mul(p, n)
	}
	sum, err := SumDigits(p.String())
	if err != nil {
		panic(err)
	}
	fmt.Println("product:", p.String(), "sum:", sum)
}

func SumDigits(input string) (int, error) {
	l := len(input)
	sum := 0
	for i := 0; i < l; i++ {
		n, err := strconv.Atoi(input[i : i+1])
		if err != nil {
			return 0, fmt.Errorf("invalid input: %s", input)
		}
		sum += n
	}
	return sum, nil
}
