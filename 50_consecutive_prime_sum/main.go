package main

import "fmt"

func main() {
	solution()
}

func solution() {
	limit := 1000000
	notPrime := make([]int, limit)
	var primes []int
	notPrime[0] = 1
	notPrime[1] = 1

	for i := 2; i < limit; i++ {
		if notPrime[i] == 0 {
			primes = append(primes, i)
			for j := i * i; j < limit; j += i {
				notPrime[j] = 1
			}
		}
	}

	maxSum := 0
	maxRun := -1
	for i := 0; i < len(primes); i++ {
		sum := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			if sum >= limit {
				break
			}
			if notPrime[sum] == 0 && sum > maxSum && j-i > maxRun {
				maxRun = j - i
				maxSum = sum
			}
		}
	}
	fmt.Println("max sum:", maxSum)
}
