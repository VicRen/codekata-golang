package main

import "fmt"

func main() {
	var nums []int
	sum := 0
	for i := 1; i < 10000; i++ {
		if isAmicableNumber(i) {
			nums = append(nums, i)
			sum += i
		}
	}
	fmt.Println("sum:", sum, "nums:", nums)
}

func sumOfDivisors(input int) int {
	sum := 0
	for i := 1; i <= input/2; i++ {
		if input%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAmicableNumber(input int) bool {
	da := sumOfDivisors(input)
	db := sumOfDivisors(da)
	if da == db {
		return false
	}
	return input == db
}
