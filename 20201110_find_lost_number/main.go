package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(100)
	fmt.Println("r:", r)
	var src []int
	for i := 0; i < 100; i++ {
		if i == r {
			continue
		}
		src = append(src, i)
	}
	fmt.Println("lost number:", findLostNumber(src))
	fmt.Println("lost number2:", findLostNumber2(src))
}

func findLostNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	n := len(nums)
	return n*(n+1)/2 - sum
}

func findLostNumber2(nums []int) int {
	res := 0
	for i, n := range nums {
		res = res ^ i ^ n
	}
	res = res ^ len(nums)
	return res
}
