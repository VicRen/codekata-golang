package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
)

func main() {
	fi, err := os.Open("./digits.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var input []*big.Int
	for {
		a, _, e := br.ReadLine()
		if e == io.EOF {
			break
		}
		i, ok := (&big.Int{}).SetString(string(a), 0)
		if !ok {
			panic("invalid number")
		}
		input = append(input, i)
	}

	sum := &big.Int{}
	for _, v := range input {
		sum = sum.Add(sum, v)
	}

	fmt.Println(sum)

	fmt.Println("ten digits is", sum.String()[:10])
}
