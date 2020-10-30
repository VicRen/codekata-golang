package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i < 355000; i++ {
		if sumOfFifthPowerOfDigits(i) == int64(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumOfFifthPowerOfDigits(n int) int64 {
	number := n
	sum := int64(0)
	for number > 0 {
		x := number % 10
		sum += fifthPower(x)
		number /= 10
	}
	return sum
}

func fifthPower(n int) int64 {
	p := int64(1)
	for i := 0; i < 5; i++ {
		p *= int64(n)
	}
	return p
}
