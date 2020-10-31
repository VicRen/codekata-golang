package main

import "fmt"

var coins = []int{1, 2, 5, 10, 20, 50, 100, 200}

func main() {
	fmt.Println(sumCoins(200))
}

func sumCoins(target int) int {
	t := make([]int, target+1)
	t[0] = 1
	for _, c := range coins {
		for j := c; j <= target; j++ {
			t[j] += t[j-c]
		}
	}
	return t[target]
}
