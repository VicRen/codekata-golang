package main

import "fmt"

func main() {
	n := 20
	for {
		if IsDivisibleFor1to20(n) {
			fmt.Println("evenly divisible by all of the numbers from 1 to 20:", n)
			return
		}
		n++
	}
}

func IsDivisibleFor1to20(n int) bool {
	for i := 1; i <= 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
