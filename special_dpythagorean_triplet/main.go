package main

import (
	"fmt"
)

func main() {
	got := SplitIntegerIntoThree(1000)
	product := 0
	for _, v := range got {
		if IsPythagoreanTriplet(v) {
			fmt.Println("Pythagorean triplet ", v)
			product = GetProduct(v)
			break
		}
	}
	fmt.Println("product:", product)
}

func SplitIntegerIntoThree(src int) [][]int {
	var ret [][]int
	x := 1
	for {
		left1 := src - x
		if left1 < x {
			break
		}
		y := x + 1
		for {
			var r []int
			r = append(r, x)
			left2 := left1 - y
			if left2 < y+1 {
				break
			}
			r = append(r, y)
			r = append(r, left2)
			ret = append(ret, r)
			y++
		}
		x++
	}
	return ret
}

func IsPythagoreanTriplet(ns []int) bool {
	if len(ns) != 3 {
		return false
	}
	var squares []int
	for _, v := range ns {
		squares = append(squares, v*v)
	}
	if squares[0]+squares[1] == squares[2] {
		return true
	}
	return false
}

func GetProduct(ns []int) int {
	product := 1
	for _, v := range ns {
		product = product * v
	}
	return product
}
