package main

func PrimeFactorsOf(n int) []int {
	var ret = make([]int, 0)
	var divider = 2
	for divider < n {
		for n % divider == 0 {
			ret = append(ret, divider)
			n = n/divider
		}
		divider++
	}
	if n > 1 {
		ret = append(ret, n)
	}
	return ret
}
