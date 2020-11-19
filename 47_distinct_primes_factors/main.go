package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := make([]int, 0)
	i := 2
	for {
		primes := distinctPrimesOf(i)
		if len(primes) == 4 {
			l := len(nums)
			if l == 0 || (nums[l-1] == i-1) {
				nums = append(nums, i)
			} else {
				nums = []int{i}
			}
			if len(nums) == 4 {
				break
			}
		}
		i++
	}
	fmt.Println(nums)
}

func distinctPrimesOf(n int) []int {
	m := make(map[int]struct{})
	max := int(math.Sqrt(float64(n)))
	divider := 2
	for divider <= max {
		if n%divider == 0 {
			if isPrime(divider) {
				m[divider] = struct{}{}
			}
			x := n / divider
			if isPrime(x) {
				m[x] = struct{}{}
			}
		}
		divider++
	}
	var ret []int
	for k := range m {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
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
