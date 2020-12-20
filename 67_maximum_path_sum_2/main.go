package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.Open("./triangle.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var input [][]int
	for {
		a, _, e := br.ReadLine()
		if e == io.EOF {
			break
		}
		got := strings.Split(string(a), " ")
		var line []int
		for _, v := range got {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			line = append(line, n)
		}
		input = append(input, line)
	}
	fmt.Println("input=", input)

	fmt.Println("maximum path sum is:", LargestSumOfTriangle(input))
}

func LargestSumOfTriangle(input [][]int) int {
	rest := input
	l := len(input)
	for l > 1 {
		hl := l - 2
		var line []int
		for n, v := range input[l-2] {
			x := v + max(input[l-1][n], input[l-1][n+1])
			line = append(line, x)
		}
		rest = rest[:hl]
		rest = append(rest, line)
		l = len(rest)
	}
	return rest[0][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
