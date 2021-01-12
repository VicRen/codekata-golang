package main

import (
	"errors"
	"fmt"
)

func main() {
	n := 26239
	for {
		fmt.Println(n)
		nums, err := findSprialDiagonals(n)
		if err != nil {
			fmt.Println(err)
			break
		}
		if primeRatio(nums) < 0.1 {
			fmt.Println(n)
			break
		}
		n += 2
	}
}

func findSprialDiagonals(width int) ([]int, error) {
	if width%2 == 0 {
		return nil, errors.New("invalid width")
	}
	x := 3
	ret := []int{1}
	cw := 3
	d := 2
	for width/cw > 0 {
		ret = append(ret, []int{x, x + d, x + 2*d, x + 3*d}...)
		x = x + 3*d + cw + 1
		cw += 2
		d = cw - 1
	}
	return ret, nil
}

func primeRatio(nums []int) float64 {
	pc := float64(0)
	for _, n := range nums {
		if isPrime(n) {
			pc++
		}
	}
	return pc / float64(len(nums))
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
