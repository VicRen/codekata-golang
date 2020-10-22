package main

import "fmt"

func main() {
	n := 20
	for {
		if IsDivisible(1, 20, n) {
			fmt.Println("evenly divisible by all of the numbers from 1 to 20:", n)
			return
		}
		n++
	}
}

func IsDivisible(start, end, n int) bool {
	for i := start; i <= end; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
