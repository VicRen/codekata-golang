package main

import (
	"math/big"
	"strconv"
)

func main() {

}

func solution(k int, a []int) []int {
	ks, _ := (&big.Int{}).SetString(strconv.Itoa(k), 0)
	s := (&big.Int{}).Add(mergeSlice(a), ks).String()
	return stringToSlice(s)
}

func mergeSlice(n []int) *big.Int {
	var s string
	for _, x := range n {
		sn := strconv.Itoa(x)
		s += sn
	}
	ret, _ := (&big.Int{}).SetString(s, 0)
	return ret
}

func stringToSlice(s string) []int {
	var ret []int
	for _, sn := range s {
		in, _ := strconv.Atoi(string(sn))
		ret = append(ret, in)
	}
	return ret
}
