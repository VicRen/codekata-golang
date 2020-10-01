package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		fb := FizzBuzz(i)
		if fb == "" {
			fmt.Println(i)
			continue
		} else {
			sum += i
		}
		fmt.Println(fb)
	}
	fmt.Println("sum=", sum)
}

func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return ""
}
