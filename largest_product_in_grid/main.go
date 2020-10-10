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
	fi, err := os.Open("./digits.txt")
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

	d4 := GetDown4(input)
	fmt.Println(d4)

	var product *Result
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			p1 := NewResult()
			p2 := NewResult()
			p3 := NewResult()
			p4 := NewResult()
			for i := 0; i < 4; i++ {
				p1.Product *= GetAdjacent(input, y, x+i)
				p1.Line = append(p1.Line, GetAdjacent(input, y, x+i))
				p2.Product *= GetAdjacent(input, y+i, x)
				p2.Line = append(p2.Line, GetAdjacent(input, y+i, x))
				p3.Product *= GetAdjacent(input, y+i, x+i)
				p3.Line = append(p3.Line, GetAdjacent(input, y+i, x+i))
				p4.Product *= GetAdjacent(input, y+i, x-i)
				p4.Line = append(p4.Line, GetAdjacent(input, y+i, x-i))
			}
			product = Max(p1, p2, p3, p4, product)
		}
	}

	fmt.Println("product is", product)
}

type Result struct {
	Product int
	Line    []int
}

func NewResult() *Result {
	return &Result{
		Product: 1,
	}
}

func (r *Result) String() string {
	return fmt.Sprintln("product is", r.Product, "line:", r.Line)
}

func Max(n, n2, n3, n4, n5 *Result) *Result {
	max := NewResult()
	if n.Product > max.Product {
		max = n
	}
	if n2.Product > max.Product {
		max = n2
	}
	if n3.Product > max.Product {
		max = n3
	}
	if n4.Product > max.Product {
		max = n4
	}
	if n5 != nil && n5.Product > max.Product {
		max = n5
	}
	return max
}

func GetAdjacent(input [][]int, x, y int) int {
	if 0 <= y && y < len(input) && 0 <= x && x < len(input[y]) {
		return input[y][x]
	}
	return 0
}

func GetRight4(input [][]int) [][]int {
	var ret [][]int
	for _, line := range input {
		l := len(line)
		if l < 4 {
			continue
		} else if l == 4 {
			ret = append(ret, line)
		} else {
			for x := 0; x+4 <= l; x++ {
				ret = append(ret, line[x:x+4])
			}
		}
	}
	return ret
}

func GetDown4(input [][]int) [][]int {
	length := len(input)
	if length < 4 {
		return nil
	}
	maxL := 0
	for _, line := range input {
		if len(line) > maxL {
			maxL = len(line)
		}
	}
	var ret [][]int
	for c := 0; c < maxL; c++ {
		for l := 0; l <= length-4; l++ {
			var one []int
			for d := 0; d < 4; d++ {
				if len(input[l+d]) > c {
					one = append(one, input[l+d][c])
				}
			}
			if len(one) == 4 {
				ret = append(ret, one)
			}
		}
	}
	return ret
}

func GetDiagonal4(input [][]int) [][]int {
	l := len(input)
	if l < 4 {
		return nil
	}
	maxL := 0
	for _, line := range input {
		if len(line) > maxL {
			maxL = len(line)
		}
	}
	if maxL < 4 {
		return nil
	}
	var ret [][]int
	for x := 0; x < maxL; x++ {
		for y := 0; y < l; y++ {
			if l < y+4 {
				break
			}
			var l []int
			for d := 0; d < 4; d++ {
				line := input[d+y]
				if len(line) <= d+x {
					break
				}
				l = append(l, line[x+d])
			}
			if len(l) == 4 {
				ret = append(ret, l)
			}
		}
	}
	return ret
}
