package main

import (
	"fmt"
	"strconv"
)

func main() {
	saved := map[int]struct{}{}
	for i := 2; i < 1e6; i++ {
		fmt.Println("i:", i)
		if _, ok := saved[i]; ok {
			continue
		}
		ns := circularNumbersOf(i)
		tmpMap := map[int]struct{}{}
		for _, n := range ns {
			if isPrime(n) {
				tmpMap[n] = struct{}{}
			}
		}
		if len(tmpMap) == len(ns) {
			for k := range tmpMap {
				saved[k] = struct{}{}
			}
		}
	}
	fmt.Println("num:", len(saved), saved)
}

func circularNumbersOf(n int) []int {
	var ret []int
	retMap := map[int]struct{}{}
	ns := strconv.Itoa(n)
	ret = append(ret, n)
	retMap[n] = struct{}{}
	l := len(ns)
	if l < 2 {
		return ret
	}
	for i := 1; i < l; i++ {
		ns = ns[1:] + ns[:1]
		num, _ := strconv.Atoi(ns)
		if _, ok := retMap[num]; ok {
			continue
		}
		retMap[num] = struct{}{}
		ret = append(ret, num)
	}
	return ret
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	res := n
	//牛顿法求平方根
	for res*res > n {
		res = (res + n/res) / 2
	}
	divider := 2
	for divider <= res {
		if n%divider == 0 {
			return false
		}
		divider++
	}
	return true
}
