package main

import "fmt"

func main() {
	sum := 0
	var nums []int
	ans := make(map[int]struct{})
	for i := 1; i < 28123; i++ {
		if isAbundantNumber(i) {
			ans[i] = struct{}{}
		}
		find := false
		for k := range ans {
			d := i - k
			if d < 0 {
				break
			}
			if _, f := ans[d]; f {
				fmt.Printf("%d = %d + %d\n", i, k, d)
				find = true
				break
			}
		}
		if !find {
			sum += i
			nums = append(nums, i)
		}
	}
	fmt.Println("abundant numbers:", ans)
	fmt.Println("sum:", sum, "nums:", nums)
}

func isAbundantNumber(input int) bool {
	return sumOfDivisors(input) > input
}

func sumOfDivisors(input int) int {
	if input < 2 {
		return 0
	}
	sum := 0
	i := 1
	for i*i <= input {
		if input%i == 0 {
			sum += i
			j := input / i
			if j != i && j != input {
				sum += j
			}
		}
		i++
	}
	return sum
}
