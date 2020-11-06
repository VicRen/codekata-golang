package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for i := 0; i < 1e6; i++ {
		if isPalindromeDecimal(i) && isPalindromeString(convertToBinary(i)) {
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}

func isPalindromeDecimal(n int) bool {
	return isPalindromeString(strconv.Itoa(n))
}

func convertToBinary(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}

func isPalindromeString(s string) bool {
	pb := strings.Builder{}
	for i := len(s); i > 0; i-- {
		pb.WriteString(s[i-1 : i])
	}
	pstr := pb.String()
	return s == pstr
}
