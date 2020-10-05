package main

import "fmt"

func main() {
	fmt.Println("sum square difference of 100 is", SquareSum(100)-SumOfSquares(100))
}

func SquareSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}
