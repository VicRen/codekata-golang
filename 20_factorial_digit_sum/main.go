package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("sum:", sumDigits(findFactorial(100)))
}

func findFactorial(input int) string {
	product, _ := (&big.Int{}).SetString("1", 0)
	for input > 0 {
		i, _ := (&big.Int{}).SetString(strconv.Itoa(input), 0)
		product = product.Mul(product, i)
		input--
	}
	return product.String()
}

func sumDigits(input string) int {
	l := len(input)
	if l == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < l; i++ {
		ns := input[i : i+1]
		n, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
		sum += n
	}
	return sum
}
