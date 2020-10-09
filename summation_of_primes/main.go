package main

import "fmt"

func main() {
	sum := 0
	from := 2
	count := 0
	for {
		if from == 2000000 {
			break
		}
		if IsPrime(from) {
			fmt.Println("prime:", from)
			sum += from
			count++
		}
		from++
	}
	fmt.Println("sum:", sum, count)
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	squareRoot := 1
	for squareRoot*squareRoot < n {
		squareRoot++
	}
	for i := 2; i <= squareRoot; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
