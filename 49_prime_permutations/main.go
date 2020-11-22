package main

import "fmt"

func main() {
	for a := 1489; a < 3340; a++ {
		b := a + 3330
		c := b + 3330
		if isPrime(a) && isPrime(b) && isPrime(c) && isPerm(a, b) && isPerm(b, c) {
			fmt.Println((a*10000+b)*10000 + c)
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	res := n
	//牛顿法求平方根
	for res*res > n {
		res = (res + n/res) / 2
	}
	divider := 2
	for divider <= res {
		if n%divider == 0 {
			return false
		}
		divider++
	}
	return true
}

func isPerm(a, b int) bool {
	cnt := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for a > 0 {
		cnt[a%10]++
		a = a / 10
	}
	for b > 0 {
		cnt[b%10]--
		b = b / 10
	}
	for _, n := range cnt {
		if n != 0 {
			return false
		}
	}
	return true
}
