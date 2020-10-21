package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	findMaxPalindrome(100, 999)
}

func findMaxPalindrome(start, end int) int {
	var n int
	var x int
	var y int
	for i := start; i < end+1; i++ {
		for j := start; j < end+1; j++ {
			tmp := i * j
			if tmp > n && IsPalindrome(tmp) {
				x = i
				y = j
				n = tmp
			}
		}
	}
	fmt.Println(x, y, n)
	return n
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
