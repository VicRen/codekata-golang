package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const BASE = 64

func main() {
	b, err := ioutil.ReadFile("p042_words.txt")
	if err != nil {
		panic(err)
	}

	word := readWords(string(b))
	count := 0
	for _, w := range word {
		if isTriangleNumber(worthOf(w)) {
			count++
		}
	}
	fmt.Println("count:", count)
}

func readWords(s string) []string {
	if len(s) == 0 {
		return nil
	}
	str := strings.Replace(s, "\"", "", -1)
	ret := strings.Split(str, ",")
	return ret
}

func worthOf(word string) int {
	if len(word) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(word); i++ {
		a := word[i : i+1][0]
		w := uint8(a) - BASE
		sum += int(w)
	}
	return sum
}

func isTriangleNumber(n int) bool {
	//tn = ½n(n+1)
	x := 2 * n
	res := x
	//牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	for i := res; i > 0; i-- {
		if i*(i+1) == x {
			return true
		}
	}
	return false
}
