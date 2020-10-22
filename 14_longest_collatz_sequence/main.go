package main

import "fmt"

func main() {
	num := 0
	ml := 0
	for i := 1000000; i > 1; i-- {
		n := i
		count := 1
		for {
			n = CollatzNext(n)
			if n == 1 {
				break
			}
			count++
		}
		if count > ml {
			ml = count
			num = i
		}
	}
	fmt.Println("largest collatz is", num, "len:", ml)
}

func CollatzNext(n int) int {
	if n == 1 {
		return 1
	} else if IsEven(n) {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func IsEven(n int) bool {
	if n > 1 && n%2 == 0 {
		return true
	}
	return false
}
