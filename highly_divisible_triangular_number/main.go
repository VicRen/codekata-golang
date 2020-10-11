package main

import "fmt"

func main() {
	n := 1
	d := 1
	for Tau(d) <= 500 {
		n++
		d += n
	}

	fmt.Println("num is", d, n)
}

func Tau(num int) int {
	if num == 1 {
		return 1
	}
	n := num
	i := 2
	p := 1

	for i*i < n {
		c := 1
		for n%i == 0 {
			n /= i
			c++
		}
		i++
		p *= c
	}

	if n == num || n > 1 {
		p *= 1 + 1
	}

	return p
}

func GetTriangularNumber(n int) int {
	return n * (n + 1) / 2
}

func ListDivisors(input int) []int {
	var ret []int
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}
