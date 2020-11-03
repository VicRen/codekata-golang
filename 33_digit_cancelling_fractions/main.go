package main

import "fmt"

func main() {
	solution()
}

func solution() {
	dp := 1
	np := 1
	for c := 1; c <= 9; c++ {
		for d := 1; d < c; d++ {
			for n := 1; n < d; n++ {
				if (n*10+c)*d == (c*10+d)*n {
					fmt.Println("d:", d, "n:", n)
					dp *= d
					np *= n
				}
			}
		}
	}
	fmt.Println("np:", np, "dp:", dp)
	fmt.Println("num:", dp/findGCD(np, dp))
}

func findGCD(a, b int) int {
	for a > 0 {
		t := a
		a = b % a
		b = t
	}
	return b
}
