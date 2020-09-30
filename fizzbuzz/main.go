package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fb := FizzBuzz(i)
		if fb == "" {
			fmt.Println(i)
			continue
		}
		fmt.Println(fb)
	}
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
