package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := 0
	for x := 9487; x >= 9234; x-- {
		res := 100002 * x
		if isPandigital(strconv.Itoa(res)) {
			num = res
			break
		}
	}
	fmt.Println("num:", num)
}

func isPandigital(s string) bool {
	var ss []string
	for _, c := range s {
		ss = append(ss, string(c))
	}
	sort.Strings(ss)
	s = ""
	for _, c := range ss {
		s += c
	}
	return s == "123456789"
}
