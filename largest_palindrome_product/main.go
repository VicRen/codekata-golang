package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	var x int
	var y int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			tmp := i * j
			if tmp > n && IsPalindrome(tmp) {
				x = i
				y = j
				n = tmp
			}
		}
	}
	fmt.Println(x, y, n)
}

func IsPalindrome(n int) bool {
	str := strconv.Itoa(n)
	pb := strings.Builder{}
	for i := len(str); i > 0; i-- {
		pb.WriteString(str[i-1 : i])
	}
	pstr := pb.String()
	return str == pstr
}
